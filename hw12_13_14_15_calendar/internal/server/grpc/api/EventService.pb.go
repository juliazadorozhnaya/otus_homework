// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: internal/server/grpc/EventService.proto

package api

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_grpc_EventService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_grpc_EventService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_internal_server_grpc_EventService_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Age       int64  `protobuf:"varint,5,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_grpc_EventService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_grpc_EventService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_internal_server_grpc_EventService_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	BeginningT    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=BeginningT,proto3" json:"BeginningT,omitempty"`
	FinishT       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=FinishT,proto3" json:"FinishT,omitempty"`
	NotificationT *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=NotificationT,proto3" json:"NotificationT,omitempty"`
	UserID        string                 `protobuf:"bytes,7,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_grpc_EventService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_grpc_EventService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_internal_server_grpc_EventService_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetBeginningT() *timestamppb.Timestamp {
	if x != nil {
		return x.BeginningT
	}
	return nil
}

func (x *Event) GetFinishT() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishT
	}
	return nil
}

func (x *Event) GetNotificationT() *timestamppb.Timestamp {
	if x != nil {
		return x.NotificationT
	}
	return nil
}

func (x *Event) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type DateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
}

func (x *DateRequest) Reset() {
	*x = DateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_grpc_EventService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRequest) ProtoMessage() {}

func (x *DateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_grpc_EventService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRequest.ProtoReflect.Descriptor instead.
func (*DateRequest) Descriptor() ([]byte, []int) {
	return file_internal_server_grpc_EventService_proto_rawDescGZIP(), []int{3}
}

func (x *DateRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_grpc_EventService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_grpc_EventService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_internal_server_grpc_EventService_proto_rawDescGZIP(), []int{4}
}

func (x *Events) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_server_grpc_EventService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_internal_server_grpc_EventService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_internal_server_grpc_EventService_proto_rawDescGZIP(), []int{5}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_internal_server_grpc_EventService_proto protoreflect.FileDescriptor

var file_internal_server_grpc_EventService_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x78, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x22, 0x9b, 0x02, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x0a, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x12, 0x34, 0x0a, 0x07, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x54, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54,
	0x12, 0x40, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3d, 0x0a, 0x0b, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x22, 0x28, 0x0a, 0x06, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xa0, 0x02, 0x0a, 0x0c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x05, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x1a, 0x07, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x12,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x44,
	0x61, 0x79, 0x12, 0x0c, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x07, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x13, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x57, 0x65,
	0x65, 0x6b, 0x12, 0x0c, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x07, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x14, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x0c, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x32, 0x69, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x05, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x1a, 0x06, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x05,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_server_grpc_EventService_proto_rawDescOnce sync.Once
	file_internal_server_grpc_EventService_proto_rawDescData = file_internal_server_grpc_EventService_proto_rawDesc
)

func file_internal_server_grpc_EventService_proto_rawDescGZIP() []byte {
	file_internal_server_grpc_EventService_proto_rawDescOnce.Do(func() {
		file_internal_server_grpc_EventService_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_server_grpc_EventService_proto_rawDescData)
	})
	return file_internal_server_grpc_EventService_proto_rawDescData
}

var (
	file_internal_server_grpc_EventService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
	file_internal_server_grpc_EventService_proto_goTypes  = []interface{}{
		(*Void)(nil),                  // 0: Void
		(*User)(nil),                  // 1: User
		(*Event)(nil),                 // 2: Event
		(*DateRequest)(nil),           // 3: DateRequest
		(*Events)(nil),                // 4: Events
		(*Users)(nil),                 // 5: Users
		(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	}
)

var file_internal_server_grpc_EventService_proto_depIdxs = []int32{
	6,  // 0: Event.BeginningT:type_name -> google.protobuf.Timestamp
	6,  // 1: Event.FinishT:type_name -> google.protobuf.Timestamp
	6,  // 2: Event.NotificationT:type_name -> google.protobuf.Timestamp
	6,  // 3: DateRequest.Date:type_name -> google.protobuf.Timestamp
	2,  // 4: Events.events:type_name -> Event
	1,  // 5: Users.users:type_name -> User
	0,  // 6: EventService.SelectEvents:input_type -> Void
	2,  // 7: EventService.CreateEvent:input_type -> Event
	2,  // 8: EventService.UpdateEvent:input_type -> Event
	2,  // 9: EventService.DeleteEvent:input_type -> Event
	3,  // 10: EventService.SelectEventsForDay:input_type -> DateRequest
	3,  // 11: EventService.SelectEventsForWeek:input_type -> DateRequest
	3,  // 12: EventService.SelectEventsForMonth:input_type -> DateRequest
	0,  // 13: UserService.SelectUsers:input_type -> Void
	1,  // 14: UserService.CreateUser:input_type -> User
	1,  // 15: UserService.DeleteUser:input_type -> User
	4,  // 16: EventService.SelectEvents:output_type -> Events
	0,  // 17: EventService.CreateEvent:output_type -> Void
	0,  // 18: EventService.UpdateEvent:output_type -> Void
	0,  // 19: EventService.DeleteEvent:output_type -> Void
	4,  // 20: EventService.SelectEventsForDay:output_type -> Events
	4,  // 21: EventService.SelectEventsForWeek:output_type -> Events
	4,  // 22: EventService.SelectEventsForMonth:output_type -> Events
	5,  // 23: UserService.SelectUsers:output_type -> Users
	0,  // 24: UserService.CreateUser:output_type -> Void
	0,  // 25: UserService.DeleteUser:output_type -> Void
	16, // [16:26] is the sub-list for method output_type
	6,  // [6:16] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_internal_server_grpc_EventService_proto_init() }
func file_internal_server_grpc_EventService_proto_init() {
	if File_internal_server_grpc_EventService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_server_grpc_EventService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_server_grpc_EventService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_server_grpc_EventService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_server_grpc_EventService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_server_grpc_EventService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Events); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_server_grpc_EventService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_server_grpc_EventService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_internal_server_grpc_EventService_proto_goTypes,
		DependencyIndexes: file_internal_server_grpc_EventService_proto_depIdxs,
		MessageInfos:      file_internal_server_grpc_EventService_proto_msgTypes,
	}.Build()
	File_internal_server_grpc_EventService_proto = out.File
	file_internal_server_grpc_EventService_proto_rawDesc = nil
	file_internal_server_grpc_EventService_proto_goTypes = nil
	file_internal_server_grpc_EventService_proto_depIdxs = nil
}
