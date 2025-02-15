// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: wso2/discovery/config/enforcer/auth_header.proto

package enforcer

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Listener and client certificate store model
type AuthHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enable outbound auth header
	EnableOutboundAuthHeader bool `protobuf:"varint,1,opt,name=enableOutboundAuthHeader,proto3" json:"enableOutboundAuthHeader,omitempty"`
	// Auth header name
	AuthorizationHeader string `protobuf:"bytes,2,opt,name=authorizationHeader,proto3" json:"authorizationHeader,omitempty"`
}

func (x *AuthHeader) Reset() {
	*x = AuthHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_config_enforcer_auth_header_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthHeader) ProtoMessage() {}

func (x *AuthHeader) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_config_enforcer_auth_header_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthHeader.ProtoReflect.Descriptor instead.
func (*AuthHeader) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_config_enforcer_auth_header_proto_rawDescGZIP(), []int{0}
}

func (x *AuthHeader) GetEnableOutboundAuthHeader() bool {
	if x != nil {
		return x.EnableOutboundAuthHeader
	}
	return false
}

func (x *AuthHeader) GetAuthorizationHeader() string {
	if x != nil {
		return x.AuthorizationHeader
	}
	return ""
}

var File_wso2_discovery_config_enforcer_auth_header_proto protoreflect.FileDescriptor

var file_wso2_discovery_config_enforcer_auth_header_proto_rawDesc = []byte{
	0x0a, 0x30, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x22, 0x7a, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x3a, 0x0a, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x13,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x8f,
	0x01, 0x0a, 0x2a, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x42, 0x0f, 0x41,
	0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x3b, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_config_enforcer_auth_header_proto_rawDescOnce sync.Once
	file_wso2_discovery_config_enforcer_auth_header_proto_rawDescData = file_wso2_discovery_config_enforcer_auth_header_proto_rawDesc
)

func file_wso2_discovery_config_enforcer_auth_header_proto_rawDescGZIP() []byte {
	file_wso2_discovery_config_enforcer_auth_header_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_config_enforcer_auth_header_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_config_enforcer_auth_header_proto_rawDescData)
	})
	return file_wso2_discovery_config_enforcer_auth_header_proto_rawDescData
}

var file_wso2_discovery_config_enforcer_auth_header_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_config_enforcer_auth_header_proto_goTypes = []interface{}{
	(*AuthHeader)(nil), // 0: wso2.discovery.config.enforcer.AuthHeader
}
var file_wso2_discovery_config_enforcer_auth_header_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wso2_discovery_config_enforcer_auth_header_proto_init() }
func file_wso2_discovery_config_enforcer_auth_header_proto_init() {
	if File_wso2_discovery_config_enforcer_auth_header_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_config_enforcer_auth_header_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthHeader); i {
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
			RawDescriptor: file_wso2_discovery_config_enforcer_auth_header_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_config_enforcer_auth_header_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_config_enforcer_auth_header_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_config_enforcer_auth_header_proto_msgTypes,
	}.Build()
	File_wso2_discovery_config_enforcer_auth_header_proto = out.File
	file_wso2_discovery_config_enforcer_auth_header_proto_rawDesc = nil
	file_wso2_discovery_config_enforcer_auth_header_proto_goTypes = nil
	file_wso2_discovery_config_enforcer_auth_header_proto_depIdxs = nil
}
