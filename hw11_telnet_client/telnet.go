package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &TClient{address: address, timeout: timeout, in: in, out: out}
}

type TClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer

	handle net.Conn
}

func (tc *TClient) Connect() error {
	h, err := net.DialTimeout("tcp", tc.address, tc.timeout)
	if err != nil {
		return err
	}
	tc.handle = h
	return nil
}

func (tc *TClient) Close() error {
	return tc.handle.Close()
}

func (tc *TClient) Send() error {
	_, err := io.Copy(tc.handle, tc.in)
	if err != nil {
		return fmt.Errorf("error sending: %w", err)
	}
	return nil
}

func (tc *TClient) Receive() error {
	_, err := io.Copy(tc.out, tc.handle)
	if err != nil {
		return fmt.Errorf("error receive: %w", err)
	}
	return err
}
