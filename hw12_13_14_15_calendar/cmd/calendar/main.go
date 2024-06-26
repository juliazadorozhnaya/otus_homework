package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/juliazadorozhnaya/otus_homework/hw12_13_14_15_calendar/internal/app"
	"github.com/juliazadorozhnaya/otus_homework/hw12_13_14_15_calendar/internal/config"
	"github.com/juliazadorozhnaya/otus_homework/hw12_13_14_15_calendar/internal/logger"
	servergrpc "github.com/juliazadorozhnaya/otus_homework/hw12_13_14_15_calendar/internal/server/grpc"
	serverhttp "github.com/juliazadorozhnaya/otus_homework/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/juliazadorozhnaya/otus_homework/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/juliazadorozhnaya/otus_homework/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/pressly/goose/v3"
)

var (
	configPath  string
	storageType string

	ErrorInvalidStorageType = errors.New("invalid storage type")
)

func init() {
	defaultConfigPath := path.Join("config", "calendar_config.toml")
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to configuration file")
	flag.StringVar(&storageType, "storage", "sql", "Type of storage. Expected values: \"memory\" || \"sql\"")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	if err := config.LoadConfig(configPath); err != nil {
		fmt.Println(err)
		return
	}
	conf := config.Get()

	log := logger.New(conf.Logger)

	var storage app.Storage
	switch storageType {
	case "memory":
		storage = memorystorage.New()
	case "sql":
		dbConn := conf.Database
		connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			dbConn.UserName, dbConn.Password, dbConn.Host, dbConn.Port, dbConn.DatabaseName)

		if err := runMigrations(connString); err != nil {
			log.Error("failed to run migrations: " + err.Error())
			return
		}

		var err error
		storage, err = sqlstorage.New(connString)
		if err != nil {
			log.Error("failed to create SQL storage: " + err.Error())
			return
		}
	default:
		log.Error(ErrorInvalidStorageType.Error())
		os.Exit(1)
	}

	calendarApp := app.New(storage, *log)
	httpServer := serverhttp.NewServer(log, calendarApp, conf.HTTPServer)
	grpcServer := servergrpc.NewServer(log, calendarApp, conf.GRPCServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := httpServer.Start(); !errors.Is(err, http.ErrServerClosed) && err != nil {
			log.Error("failed to start HTTP server: " + err.Error())
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		if err := grpcServer.Start(); err != nil {
			log.Error("failed to start gRPC server: " + err.Error())
			cancel()
		}
	}()

	go func() {
		<-ctx.Done()
		log.Info("shutting down servers...")

		if err := httpServer.Stop(ctx); err != nil {
			log.Error("failed to stop HTTP server: " + err.Error())
		}

		if err := grpcServer.Stop(ctx); err != nil {
			log.Error("failed to stop GRPC server: " + err.Error())
		}
	}()

	log.Info("App is running...")
	wg.Wait()
	log.Info("Servers closed")
}

func runMigrations(connString string) error {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := goose.Up(db, "migrations", goose.WithAllowMissing(), goose.WithNoVersioning()); err != nil {
		return err
	}
	return nil
}
