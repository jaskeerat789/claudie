// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: proto/terraformer.proto

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

type BuildInfrastructureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentState *Project `protobuf:"bytes,1,opt,name=currentState,proto3" json:"currentState,omitempty"`
	DesiredState *Project `protobuf:"bytes,2,opt,name=desiredState,proto3" json:"desiredState,omitempty"`
}

func (x *BuildInfrastructureRequest) Reset() {
	*x = BuildInfrastructureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_terraformer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfrastructureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfrastructureRequest) ProtoMessage() {}

func (x *BuildInfrastructureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_terraformer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfrastructureRequest.ProtoReflect.Descriptor instead.
func (*BuildInfrastructureRequest) Descriptor() ([]byte, []int) {
	return file_proto_terraformer_proto_rawDescGZIP(), []int{0}
}

func (x *BuildInfrastructureRequest) GetCurrentState() *Project {
	if x != nil {
		return x.CurrentState
	}
	return nil
}

func (x *BuildInfrastructureRequest) GetDesiredState() *Project {
	if x != nil {
		return x.DesiredState
	}
	return nil
}

type BuildInfrastructureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentState *Project `protobuf:"bytes,1,opt,name=currentState,proto3" json:"currentState,omitempty"`
	DesiredState *Project `protobuf:"bytes,2,opt,name=desiredState,proto3" json:"desiredState,omitempty"`
	ErrorMessage string   `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *BuildInfrastructureResponse) Reset() {
	*x = BuildInfrastructureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_terraformer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfrastructureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfrastructureResponse) ProtoMessage() {}

func (x *BuildInfrastructureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_terraformer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfrastructureResponse.ProtoReflect.Descriptor instead.
func (*BuildInfrastructureResponse) Descriptor() ([]byte, []int) {
	return file_proto_terraformer_proto_rawDescGZIP(), []int{1}
}

func (x *BuildInfrastructureResponse) GetCurrentState() *Project {
	if x != nil {
		return x.CurrentState
	}
	return nil
}

func (x *BuildInfrastructureResponse) GetDesiredState() *Project {
	if x != nil {
		return x.DesiredState
	}
	return nil
}

func (x *BuildInfrastructureResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type DestroyInfrastructureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *DestroyInfrastructureRequest) Reset() {
	*x = DestroyInfrastructureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_terraformer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyInfrastructureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyInfrastructureRequest) ProtoMessage() {}

func (x *DestroyInfrastructureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_terraformer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyInfrastructureRequest.ProtoReflect.Descriptor instead.
func (*DestroyInfrastructureRequest) Descriptor() ([]byte, []int) {
	return file_proto_terraformer_proto_rawDescGZIP(), []int{2}
}

func (x *DestroyInfrastructureRequest) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type DestroyInfrastructureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *DestroyInfrastructureResponse) Reset() {
	*x = DestroyInfrastructureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_terraformer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyInfrastructureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyInfrastructureResponse) ProtoMessage() {}

func (x *DestroyInfrastructureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_terraformer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyInfrastructureResponse.ProtoReflect.Descriptor instead.
func (*DestroyInfrastructureResponse) Descriptor() ([]byte, []int) {
	return file_proto_terraformer_proto_rawDescGZIP(), []int{3}
}

func (x *DestroyInfrastructureResponse) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_proto_terraformer_proto protoreflect.FileDescriptor

var file_proto_terraformer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6c, 0x61, 0x75, 0x64,
	0x69, 0x65, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x1a, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6c,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0xad, 0x01, 0x0a, 0x1b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x72,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x47, 0x0a, 0x1c, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x48, 0x0a, 0x1d, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c,
	0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x32, 0xde, 0x01, 0x0a, 0x12, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a,
	0x15, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x49,
	0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_terraformer_proto_rawDescOnce sync.Once
	file_proto_terraformer_proto_rawDescData = file_proto_terraformer_proto_rawDesc
)

func file_proto_terraformer_proto_rawDescGZIP() []byte {
	file_proto_terraformer_proto_rawDescOnce.Do(func() {
		file_proto_terraformer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_terraformer_proto_rawDescData)
	})
	return file_proto_terraformer_proto_rawDescData
}

var file_proto_terraformer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_terraformer_proto_goTypes = []interface{}{
	(*BuildInfrastructureRequest)(nil),    // 0: claudie.BuildInfrastructureRequest
	(*BuildInfrastructureResponse)(nil),   // 1: claudie.BuildInfrastructureResponse
	(*DestroyInfrastructureRequest)(nil),  // 2: claudie.DestroyInfrastructureRequest
	(*DestroyInfrastructureResponse)(nil), // 3: claudie.DestroyInfrastructureResponse
	(*Project)(nil),                       // 4: claudie.Project
	(*Config)(nil),                        // 5: claudie.Config
}
var file_proto_terraformer_proto_depIdxs = []int32{
	4, // 0: claudie.BuildInfrastructureRequest.currentState:type_name -> claudie.Project
	4, // 1: claudie.BuildInfrastructureRequest.desiredState:type_name -> claudie.Project
	4, // 2: claudie.BuildInfrastructureResponse.currentState:type_name -> claudie.Project
	4, // 3: claudie.BuildInfrastructureResponse.desiredState:type_name -> claudie.Project
	5, // 4: claudie.DestroyInfrastructureRequest.config:type_name -> claudie.Config
	5, // 5: claudie.DestroyInfrastructureResponse.config:type_name -> claudie.Config
	0, // 6: claudie.TerraformerService.BuildInfrastructure:input_type -> claudie.BuildInfrastructureRequest
	2, // 7: claudie.TerraformerService.DestroyInfrastructure:input_type -> claudie.DestroyInfrastructureRequest
	1, // 8: claudie.TerraformerService.BuildInfrastructure:output_type -> claudie.BuildInfrastructureResponse
	3, // 9: claudie.TerraformerService.DestroyInfrastructure:output_type -> claudie.DestroyInfrastructureResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_terraformer_proto_init() }
func file_proto_terraformer_proto_init() {
	if File_proto_terraformer_proto != nil {
		return
	}
	file_proto_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_terraformer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfrastructureRequest); i {
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
		file_proto_terraformer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfrastructureResponse); i {
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
		file_proto_terraformer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyInfrastructureRequest); i {
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
		file_proto_terraformer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyInfrastructureResponse); i {
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
			RawDescriptor: file_proto_terraformer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_terraformer_proto_goTypes,
		DependencyIndexes: file_proto_terraformer_proto_depIdxs,
		MessageInfos:      file_proto_terraformer_proto_msgTypes,
	}.Build()
	File_proto_terraformer_proto = out.File
	file_proto_terraformer_proto_rawDesc = nil
	file_proto_terraformer_proto_goTypes = nil
	file_proto_terraformer_proto_depIdxs = nil
}
