// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: pb/userInterface.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddUserInterfaceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	InterfaceInfoId string `protobuf:"bytes,2,opt,name=interfaceInfoId,proto3" json:"interfaceInfoId,omitempty"`
}

func (x *AddUserInterfaceInfoRequest) Reset() {
	*x = AddUserInterfaceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userInterface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserInterfaceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserInterfaceInfoRequest) ProtoMessage() {}

func (x *AddUserInterfaceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userInterface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserInterfaceInfoRequest.ProtoReflect.Descriptor instead.
func (*AddUserInterfaceInfoRequest) Descriptor() ([]byte, []int) {
	return file_pb_userInterface_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserInterfaceInfoRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddUserInterfaceInfoRequest) GetInterfaceInfoId() string {
	if x != nil {
		return x.InterfaceInfoId
	}
	return ""
}

type AddUserInterfaceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   bool   `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddUserInterfaceInfoResponse) Reset() {
	*x = AddUserInterfaceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userInterface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserInterfaceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserInterfaceInfoResponse) ProtoMessage() {}

func (x *AddUserInterfaceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userInterface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserInterfaceInfoResponse.ProtoReflect.Descriptor instead.
func (*AddUserInterfaceInfoResponse) Descriptor() ([]byte, []int) {
	return file_pb_userInterface_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserInterfaceInfoResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddUserInterfaceInfoResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *AddUserInterfaceInfoResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type GetUserByAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey string `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
}

func (x *GetUserByAccessKeyRequest) Reset() {
	*x = GetUserByAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userInterface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByAccessKeyRequest) ProtoMessage() {}

func (x *GetUserByAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userInterface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*GetUserByAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_pb_userInterface_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByAccessKeyRequest) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	SecretKey string `protobuf:"bytes,2,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userInterface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userInterface_proto_msgTypes[3]
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
	return file_pb_userInterface_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

type GetUserByAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *User  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserByAccessKeyResponse) Reset() {
	*x = GetUserByAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userInterface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByAccessKeyResponse) ProtoMessage() {}

func (x *GetUserByAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userInterface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*GetUserByAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_pb_userInterface_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserByAccessKeyResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetUserByAccessKeyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetUserByAccessKeyResponse) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetInterfaceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *GetInterfaceInfoRequest) Reset() {
	*x = GetInterfaceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userInterface_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInterfaceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInterfaceInfoRequest) ProtoMessage() {}

func (x *GetInterfaceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userInterface_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInterfaceInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInterfaceInfoRequest) Descriptor() ([]byte, []int) {
	return file_pb_userInterface_proto_rawDescGZIP(), []int{5}
}

func (x *GetInterfaceInfoRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetInterfaceInfoRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type GetInterfaceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetInterfaceInfoResponse) Reset() {
	*x = GetInterfaceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_userInterface_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInterfaceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInterfaceInfoResponse) ProtoMessage() {}

func (x *GetInterfaceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_userInterface_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInterfaceInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInterfaceInfoResponse) Descriptor() ([]byte, []int) {
	return file_pb_userInterface_proto_rawDescGZIP(), []int{6}
}

func (x *GetInterfaceInfoResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetInterfaceInfoResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetInterfaceInfoResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_pb_userInterface_proto protoreflect.FileDescriptor

var file_pb_userInterface_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x22, 0x5f, 0x0a, 0x1b, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x1c, 0x41, 0x64, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x22, 0x73, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x5c, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xdd, 0x02, 0x0a, 0x14,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_userInterface_proto_rawDescOnce sync.Once
	file_pb_userInterface_proto_rawDescData = file_pb_userInterface_proto_rawDesc
)

func file_pb_userInterface_proto_rawDescGZIP() []byte {
	file_pb_userInterface_proto_rawDescOnce.Do(func() {
		file_pb_userInterface_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_userInterface_proto_rawDescData)
	})
	return file_pb_userInterface_proto_rawDescData
}

var file_pb_userInterface_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_userInterface_proto_goTypes = []interface{}{
	(*AddUserInterfaceInfoRequest)(nil),  // 0: userInterface.AddUserInterfaceInfoRequest
	(*AddUserInterfaceInfoResponse)(nil), // 1: userInterface.AddUserInterfaceInfoResponse
	(*GetUserByAccessKeyRequest)(nil),    // 2: userInterface.GetUserByAccessKeyRequest
	(*User)(nil),                         // 3: userInterface.User
	(*GetUserByAccessKeyResponse)(nil),   // 4: userInterface.GetUserByAccessKeyResponse
	(*GetInterfaceInfoRequest)(nil),      // 5: userInterface.GetInterfaceInfoRequest
	(*GetInterfaceInfoResponse)(nil),     // 6: userInterface.GetInterfaceInfoResponse
}
var file_pb_userInterface_proto_depIdxs = []int32{
	3, // 0: userInterface.GetUserByAccessKeyResponse.data:type_name -> userInterface.User
	0, // 1: userInterface.UserInterfaceService.AddUserInterfaceInfo:input_type -> userInterface.AddUserInterfaceInfoRequest
	2, // 2: userInterface.UserInterfaceService.GetUserByAccessKey:input_type -> userInterface.GetUserByAccessKeyRequest
	5, // 3: userInterface.UserInterfaceService.GetInterfaceInfo:input_type -> userInterface.GetInterfaceInfoRequest
	1, // 4: userInterface.UserInterfaceService.AddUserInterfaceInfo:output_type -> userInterface.AddUserInterfaceInfoResponse
	4, // 5: userInterface.UserInterfaceService.GetUserByAccessKey:output_type -> userInterface.GetUserByAccessKeyResponse
	6, // 6: userInterface.UserInterfaceService.GetInterfaceInfo:output_type -> userInterface.GetInterfaceInfoResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_userInterface_proto_init() }
func file_pb_userInterface_proto_init() {
	if File_pb_userInterface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_userInterface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserInterfaceInfoRequest); i {
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
		file_pb_userInterface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserInterfaceInfoResponse); i {
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
		file_pb_userInterface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByAccessKeyRequest); i {
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
		file_pb_userInterface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_userInterface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByAccessKeyResponse); i {
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
		file_pb_userInterface_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInterfaceInfoRequest); i {
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
		file_pb_userInterface_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInterfaceInfoResponse); i {
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
			RawDescriptor: file_pb_userInterface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_userInterface_proto_goTypes,
		DependencyIndexes: file_pb_userInterface_proto_depIdxs,
		MessageInfos:      file_pb_userInterface_proto_msgTypes,
	}.Build()
	File_pb_userInterface_proto = out.File
	file_pb_userInterface_proto_rawDesc = nil
	file_pb_userInterface_proto_goTypes = nil
	file_pb_userInterface_proto_depIdxs = nil
}
