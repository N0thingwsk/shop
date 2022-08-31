// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: userop/v1/userop.proto

package v1

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

type CreateUseropRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUseropRequest) Reset() {
	*x = CreateUseropRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUseropRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUseropRequest) ProtoMessage() {}

func (x *CreateUseropRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUseropRequest.ProtoReflect.Descriptor instead.
func (*CreateUseropRequest) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{0}
}

type CreateUseropReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUseropReply) Reset() {
	*x = CreateUseropReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUseropReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUseropReply) ProtoMessage() {}

func (x *CreateUseropReply) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUseropReply.ProtoReflect.Descriptor instead.
func (*CreateUseropReply) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{1}
}

type UpdateUseropRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUseropRequest) Reset() {
	*x = UpdateUseropRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUseropRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUseropRequest) ProtoMessage() {}

func (x *UpdateUseropRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUseropRequest.ProtoReflect.Descriptor instead.
func (*UpdateUseropRequest) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{2}
}

type UpdateUseropReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUseropReply) Reset() {
	*x = UpdateUseropReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUseropReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUseropReply) ProtoMessage() {}

func (x *UpdateUseropReply) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUseropReply.ProtoReflect.Descriptor instead.
func (*UpdateUseropReply) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{3}
}

type DeleteUseropRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUseropRequest) Reset() {
	*x = DeleteUseropRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUseropRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUseropRequest) ProtoMessage() {}

func (x *DeleteUseropRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUseropRequest.ProtoReflect.Descriptor instead.
func (*DeleteUseropRequest) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{4}
}

type DeleteUseropReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUseropReply) Reset() {
	*x = DeleteUseropReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUseropReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUseropReply) ProtoMessage() {}

func (x *DeleteUseropReply) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUseropReply.ProtoReflect.Descriptor instead.
func (*DeleteUseropReply) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{5}
}

type GetUseropRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUseropRequest) Reset() {
	*x = GetUseropRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUseropRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUseropRequest) ProtoMessage() {}

func (x *GetUseropRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUseropRequest.ProtoReflect.Descriptor instead.
func (*GetUseropRequest) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{6}
}

type GetUseropReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUseropReply) Reset() {
	*x = GetUseropReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUseropReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUseropReply) ProtoMessage() {}

func (x *GetUseropReply) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUseropReply.ProtoReflect.Descriptor instead.
func (*GetUseropReply) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{7}
}

type ListUseropRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUseropRequest) Reset() {
	*x = ListUseropRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUseropRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUseropRequest) ProtoMessage() {}

func (x *ListUseropRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUseropRequest.ProtoReflect.Descriptor instead.
func (*ListUseropRequest) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{8}
}

type ListUseropReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUseropReply) Reset() {
	*x = ListUseropReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userop_v1_userop_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUseropReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUseropReply) ProtoMessage() {}

func (x *ListUseropReply) ProtoReflect() protoreflect.Message {
	mi := &file_userop_v1_userop_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUseropReply.ProtoReflect.Descriptor instead.
func (*ListUseropReply) Descriptor() ([]byte, []int) {
	return file_userop_v1_userop_proto_rawDescGZIP(), []int{9}
}

var File_userop_v1_userop_proto protoreflect.FileDescriptor

var file_userop_v1_userop_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa7, 0x03, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x6f, 0x70, 0x12, 0x54, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x6f, 0x70, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x54, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70,
	0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6f,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x6f, 0x70, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f,
	0x70, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x28, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x15, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userop_v1_userop_proto_rawDescOnce sync.Once
	file_userop_v1_userop_proto_rawDescData = file_userop_v1_userop_proto_rawDesc
)

func file_userop_v1_userop_proto_rawDescGZIP() []byte {
	file_userop_v1_userop_proto_rawDescOnce.Do(func() {
		file_userop_v1_userop_proto_rawDescData = protoimpl.X.CompressGZIP(file_userop_v1_userop_proto_rawDescData)
	})
	return file_userop_v1_userop_proto_rawDescData
}

var file_userop_v1_userop_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_userop_v1_userop_proto_goTypes = []interface{}{
	(*CreateUseropRequest)(nil), // 0: api.userop.v1.CreateUseropRequest
	(*CreateUseropReply)(nil),   // 1: api.userop.v1.CreateUseropReply
	(*UpdateUseropRequest)(nil), // 2: api.userop.v1.UpdateUseropRequest
	(*UpdateUseropReply)(nil),   // 3: api.userop.v1.UpdateUseropReply
	(*DeleteUseropRequest)(nil), // 4: api.userop.v1.DeleteUseropRequest
	(*DeleteUseropReply)(nil),   // 5: api.userop.v1.DeleteUseropReply
	(*GetUseropRequest)(nil),    // 6: api.userop.v1.GetUseropRequest
	(*GetUseropReply)(nil),      // 7: api.userop.v1.GetUseropReply
	(*ListUseropRequest)(nil),   // 8: api.userop.v1.ListUseropRequest
	(*ListUseropReply)(nil),     // 9: api.userop.v1.ListUseropReply
}
var file_userop_v1_userop_proto_depIdxs = []int32{
	0, // 0: api.userop.v1.Userop.CreateUserop:input_type -> api.userop.v1.CreateUseropRequest
	2, // 1: api.userop.v1.Userop.UpdateUserop:input_type -> api.userop.v1.UpdateUseropRequest
	4, // 2: api.userop.v1.Userop.DeleteUserop:input_type -> api.userop.v1.DeleteUseropRequest
	6, // 3: api.userop.v1.Userop.GetUserop:input_type -> api.userop.v1.GetUseropRequest
	8, // 4: api.userop.v1.Userop.ListUserop:input_type -> api.userop.v1.ListUseropRequest
	1, // 5: api.userop.v1.Userop.CreateUserop:output_type -> api.userop.v1.CreateUseropReply
	3, // 6: api.userop.v1.Userop.UpdateUserop:output_type -> api.userop.v1.UpdateUseropReply
	5, // 7: api.userop.v1.Userop.DeleteUserop:output_type -> api.userop.v1.DeleteUseropReply
	7, // 8: api.userop.v1.Userop.GetUserop:output_type -> api.userop.v1.GetUseropReply
	9, // 9: api.userop.v1.Userop.ListUserop:output_type -> api.userop.v1.ListUseropReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userop_v1_userop_proto_init() }
func file_userop_v1_userop_proto_init() {
	if File_userop_v1_userop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userop_v1_userop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUseropRequest); i {
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
		file_userop_v1_userop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUseropReply); i {
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
		file_userop_v1_userop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUseropRequest); i {
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
		file_userop_v1_userop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUseropReply); i {
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
		file_userop_v1_userop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUseropRequest); i {
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
		file_userop_v1_userop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUseropReply); i {
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
		file_userop_v1_userop_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUseropRequest); i {
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
		file_userop_v1_userop_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUseropReply); i {
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
		file_userop_v1_userop_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUseropRequest); i {
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
		file_userop_v1_userop_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUseropReply); i {
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
			RawDescriptor: file_userop_v1_userop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userop_v1_userop_proto_goTypes,
		DependencyIndexes: file_userop_v1_userop_proto_depIdxs,
		MessageInfos:      file_userop_v1_userop_proto_msgTypes,
	}.Build()
	File_userop_v1_userop_proto = out.File
	file_userop_v1_userop_proto_rawDesc = nil
	file_userop_v1_userop_proto_goTypes = nil
	file_userop_v1_userop_proto_depIdxs = nil
}
