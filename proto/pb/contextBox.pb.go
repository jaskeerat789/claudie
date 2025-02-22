// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: proto/contextBox.proto

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

type IdType int32

const (
	IdType_NAME IdType = 0
	IdType_HASH IdType = 1
)

// Enum value maps for IdType.
var (
	IdType_name = map[int32]string{
		0: "NAME",
		1: "HASH",
	}
	IdType_value = map[string]int32{
		"NAME": 0,
		"HASH": 1,
	}
)

func (x IdType) Enum() *IdType {
	p := new(IdType)
	*p = x
	return p
}

func (x IdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_contextBox_proto_enumTypes[0].Descriptor()
}

func (IdType) Type() protoreflect.EnumType {
	return &file_proto_contextBox_proto_enumTypes[0]
}

func (x IdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdType.Descriptor instead.
func (IdType) EnumDescriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{0}
}

// Save
type SaveConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SaveConfigRequest) Reset() {
	*x = SaveConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveConfigRequest) ProtoMessage() {}

func (x *SaveConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveConfigRequest.ProtoReflect.Descriptor instead.
func (*SaveConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{0}
}

func (x *SaveConfigRequest) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type SaveConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *SaveConfigResponse) Reset() {
	*x = SaveConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveConfigResponse) ProtoMessage() {}

func (x *SaveConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveConfigResponse.ProtoReflect.Descriptor instead.
func (*SaveConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{1}
}

func (x *SaveConfigResponse) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// Get
type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{2}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{3}
}

func (x *GetConfigResponse) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type GetAllConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllConfigsRequest) Reset() {
	*x = GetAllConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllConfigsRequest) ProtoMessage() {}

func (x *GetAllConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllConfigsRequest.ProtoReflect.Descriptor instead.
func (*GetAllConfigsRequest) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{4}
}

type GetAllConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs []*Config `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *GetAllConfigsResponse) Reset() {
	*x = GetAllConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllConfigsResponse) ProtoMessage() {}

func (x *GetAllConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllConfigsResponse.ProtoReflect.Descriptor instead.
func (*GetAllConfigsResponse) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllConfigsResponse) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

type GetConfigFromDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type IdType `protobuf:"varint,2,opt,name=type,proto3,enum=claudie.IdType" json:"type,omitempty"`
}

func (x *GetConfigFromDBRequest) Reset() {
	*x = GetConfigFromDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigFromDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigFromDBRequest) ProtoMessage() {}

func (x *GetConfigFromDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigFromDBRequest.ProtoReflect.Descriptor instead.
func (*GetConfigFromDBRequest) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{6}
}

func (x *GetConfigFromDBRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetConfigFromDBRequest) GetType() IdType {
	if x != nil {
		return x.Type
	}
	return IdType_NAME
}

type GetConfigFromDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetConfigFromDBResponse) Reset() {
	*x = GetConfigFromDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigFromDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigFromDBResponse) ProtoMessage() {}

func (x *GetConfigFromDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigFromDBResponse.ProtoReflect.Descriptor instead.
func (*GetConfigFromDBResponse) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{7}
}

func (x *GetConfigFromDBResponse) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type GetConfigByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetConfigByNameRequest) Reset() {
	*x = GetConfigByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigByNameRequest) ProtoMessage() {}

func (x *GetConfigByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigByNameRequest.ProtoReflect.Descriptor instead.
func (*GetConfigByNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{8}
}

func (x *GetConfigByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetConfigByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetConfigByNameResponse) Reset() {
	*x = GetConfigByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigByNameResponse) ProtoMessage() {}

func (x *GetConfigByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigByNameResponse.ProtoReflect.Descriptor instead.
func (*GetConfigByNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{9}
}

func (x *GetConfigByNameResponse) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

// Delete
type DeleteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type IdType `protobuf:"varint,2,opt,name=type,proto3,enum=claudie.IdType" json:"type,omitempty"`
}

func (x *DeleteConfigRequest) Reset() {
	*x = DeleteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigRequest) ProtoMessage() {}

func (x *DeleteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteConfigRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteConfigRequest) GetType() IdType {
	if x != nil {
		return x.Type
	}
	return IdType_NAME
}

type DeleteConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteConfigResponse) Reset() {
	*x = DeleteConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contextBox_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigResponse) ProtoMessage() {}

func (x *DeleteConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contextBox_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigResponse.ProtoReflect.Descriptor instead.
func (*DeleteConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_contextBox_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteConfigResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_contextBox_proto protoreflect.FileDescriptor

var file_proto_contextBox_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x3d, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x22, 0x4d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x72, 0x6f,
	0x6d, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64,
	0x69, 0x65, 0x2e, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x42, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x72, 0x6f, 0x6d,
	0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x42, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4a, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x2e, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x1c, 0x0a, 0x06, 0x49, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x41, 0x53, 0x48, 0x10, 0x01, 0x32, 0xb4, 0x06, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a,
	0x12, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x45, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x13,
	0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x11,
	0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x12, 0x1f, 0x2e,
	0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x63,
	0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x1d, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c,
	0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63,
	0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x72, 0x6f, 0x6d, 0x44,
	0x42, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_contextBox_proto_rawDescOnce sync.Once
	file_proto_contextBox_proto_rawDescData = file_proto_contextBox_proto_rawDesc
)

func file_proto_contextBox_proto_rawDescGZIP() []byte {
	file_proto_contextBox_proto_rawDescOnce.Do(func() {
		file_proto_contextBox_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_contextBox_proto_rawDescData)
	})
	return file_proto_contextBox_proto_rawDescData
}

var file_proto_contextBox_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_contextBox_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_contextBox_proto_goTypes = []interface{}{
	(IdType)(0),                     // 0: claudie.IdType
	(*SaveConfigRequest)(nil),       // 1: claudie.SaveConfigRequest
	(*SaveConfigResponse)(nil),      // 2: claudie.SaveConfigResponse
	(*GetConfigRequest)(nil),        // 3: claudie.GetConfigRequest
	(*GetConfigResponse)(nil),       // 4: claudie.GetConfigResponse
	(*GetAllConfigsRequest)(nil),    // 5: claudie.GetAllConfigsRequest
	(*GetAllConfigsResponse)(nil),   // 6: claudie.GetAllConfigsResponse
	(*GetConfigFromDBRequest)(nil),  // 7: claudie.GetConfigFromDBRequest
	(*GetConfigFromDBResponse)(nil), // 8: claudie.GetConfigFromDBResponse
	(*GetConfigByNameRequest)(nil),  // 9: claudie.GetConfigByNameRequest
	(*GetConfigByNameResponse)(nil), // 10: claudie.GetConfigByNameResponse
	(*DeleteConfigRequest)(nil),     // 11: claudie.DeleteConfigRequest
	(*DeleteConfigResponse)(nil),    // 12: claudie.DeleteConfigResponse
	(*Config)(nil),                  // 13: claudie.Config
}
var file_proto_contextBox_proto_depIdxs = []int32{
	13, // 0: claudie.SaveConfigRequest.config:type_name -> claudie.Config
	13, // 1: claudie.SaveConfigResponse.config:type_name -> claudie.Config
	13, // 2: claudie.GetConfigResponse.config:type_name -> claudie.Config
	13, // 3: claudie.GetAllConfigsResponse.configs:type_name -> claudie.Config
	0,  // 4: claudie.GetConfigFromDBRequest.type:type_name -> claudie.IdType
	13, // 5: claudie.GetConfigFromDBResponse.config:type_name -> claudie.Config
	13, // 6: claudie.GetConfigByNameResponse.config:type_name -> claudie.Config
	0,  // 7: claudie.DeleteConfigRequest.type:type_name -> claudie.IdType
	1,  // 8: claudie.ContextBoxService.SaveConfigFrontEnd:input_type -> claudie.SaveConfigRequest
	1,  // 9: claudie.ContextBoxService.SaveConfigScheduler:input_type -> claudie.SaveConfigRequest
	1,  // 10: claudie.ContextBoxService.SaveConfigBuilder:input_type -> claudie.SaveConfigRequest
	7,  // 11: claudie.ContextBoxService.GetConfigFromDB:input_type -> claudie.GetConfigFromDBRequest
	9,  // 12: claudie.ContextBoxService.GetConfigByName:input_type -> claudie.GetConfigByNameRequest
	3,  // 13: claudie.ContextBoxService.GetConfigScheduler:input_type -> claudie.GetConfigRequest
	3,  // 14: claudie.ContextBoxService.GetConfigBuilder:input_type -> claudie.GetConfigRequest
	5,  // 15: claudie.ContextBoxService.GetAllConfigs:input_type -> claudie.GetAllConfigsRequest
	11, // 16: claudie.ContextBoxService.DeleteConfig:input_type -> claudie.DeleteConfigRequest
	11, // 17: claudie.ContextBoxService.DeleteConfigFromDB:input_type -> claudie.DeleteConfigRequest
	2,  // 18: claudie.ContextBoxService.SaveConfigFrontEnd:output_type -> claudie.SaveConfigResponse
	2,  // 19: claudie.ContextBoxService.SaveConfigScheduler:output_type -> claudie.SaveConfigResponse
	2,  // 20: claudie.ContextBoxService.SaveConfigBuilder:output_type -> claudie.SaveConfigResponse
	8,  // 21: claudie.ContextBoxService.GetConfigFromDB:output_type -> claudie.GetConfigFromDBResponse
	10, // 22: claudie.ContextBoxService.GetConfigByName:output_type -> claudie.GetConfigByNameResponse
	4,  // 23: claudie.ContextBoxService.GetConfigScheduler:output_type -> claudie.GetConfigResponse
	4,  // 24: claudie.ContextBoxService.GetConfigBuilder:output_type -> claudie.GetConfigResponse
	6,  // 25: claudie.ContextBoxService.GetAllConfigs:output_type -> claudie.GetAllConfigsResponse
	12, // 26: claudie.ContextBoxService.DeleteConfig:output_type -> claudie.DeleteConfigResponse
	12, // 27: claudie.ContextBoxService.DeleteConfigFromDB:output_type -> claudie.DeleteConfigResponse
	18, // [18:28] is the sub-list for method output_type
	8,  // [8:18] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_contextBox_proto_init() }
func file_proto_contextBox_proto_init() {
	if File_proto_contextBox_proto != nil {
		return
	}
	file_proto_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_contextBox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveConfigRequest); i {
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
		file_proto_contextBox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveConfigResponse); i {
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
		file_proto_contextBox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_proto_contextBox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResponse); i {
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
		file_proto_contextBox_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllConfigsRequest); i {
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
		file_proto_contextBox_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllConfigsResponse); i {
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
		file_proto_contextBox_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigFromDBRequest); i {
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
		file_proto_contextBox_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigFromDBResponse); i {
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
		file_proto_contextBox_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigByNameRequest); i {
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
		file_proto_contextBox_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigByNameResponse); i {
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
		file_proto_contextBox_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfigRequest); i {
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
		file_proto_contextBox_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfigResponse); i {
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
			RawDescriptor: file_proto_contextBox_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_contextBox_proto_goTypes,
		DependencyIndexes: file_proto_contextBox_proto_depIdxs,
		EnumInfos:         file_proto_contextBox_proto_enumTypes,
		MessageInfos:      file_proto_contextBox_proto_msgTypes,
	}.Build()
	File_proto_contextBox_proto = out.File
	file_proto_contextBox_proto_rawDesc = nil
	file_proto_contextBox_proto_goTypes = nil
	file_proto_contextBox_proto_depIdxs = nil
}
