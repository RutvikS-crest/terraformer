// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/api/operation.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Operation is annotation for rpc that returns longrunning operation, describes
// message types that will be returned in metadata [google.protobuf.Any], and
// in response [google.protobuf.Any] (for successful operation).
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. If present, rpc returns operation which metadata field will
	// contains message of specified type.
	Metadata string `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"` // Optional.
	// Required. rpc returns operation, in case of success response will contains message of
	// specified field.
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"` // Required.
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_api_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_api_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_api_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Operation) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var file_yandex_cloud_api_operation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Operation)(nil),
		Field:         87334,
		Name:          "yandex.cloud.api.operation",
		Tag:           "bytes,87334,opt,name=operation",
		Filename:      "yandex/cloud/api/operation.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional yandex.cloud.api.Operation operation = 87334;
	E_Operation = &file_yandex_cloud_api_operation_proto_extTypes[0]
)

var File_yandex_cloud_api_operation_proto protoreflect.FileDescriptor

var file_yandex_cloud_api_operation_proto_rawDesc = []byte{
	0x0a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x5b, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa6, 0xaa, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_api_operation_proto_rawDescOnce sync.Once
	file_yandex_cloud_api_operation_proto_rawDescData = file_yandex_cloud_api_operation_proto_rawDesc
)

func file_yandex_cloud_api_operation_proto_rawDescGZIP() []byte {
	file_yandex_cloud_api_operation_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_api_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_api_operation_proto_rawDescData)
	})
	return file_yandex_cloud_api_operation_proto_rawDescData
}

var file_yandex_cloud_api_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_api_operation_proto_goTypes = []interface{}{
	(*Operation)(nil),                  // 0: yandex.cloud.api.Operation
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_yandex_cloud_api_operation_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.api.operation:extendee -> google.protobuf.MethodOptions
	0, // 1: yandex.cloud.api.operation:type_name -> yandex.cloud.api.Operation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_api_operation_proto_init() }
func file_yandex_cloud_api_operation_proto_init() {
	if File_yandex_cloud_api_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_api_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
			RawDescriptor: file_yandex_cloud_api_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_api_operation_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_api_operation_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_api_operation_proto_msgTypes,
		ExtensionInfos:    file_yandex_cloud_api_operation_proto_extTypes,
	}.Build()
	File_yandex_cloud_api_operation_proto = out.File
	file_yandex_cloud_api_operation_proto_rawDesc = nil
	file_yandex_cloud_api_operation_proto_goTypes = nil
	file_yandex_cloud_api_operation_proto_depIdxs = nil
}
