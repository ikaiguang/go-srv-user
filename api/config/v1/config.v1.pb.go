// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/config/v1/config.v1.proto

package configv1

import (
	v1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
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

// Bootstrap 配置引导
type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// app application
	App *v1.App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	// log 日志
	Log *v1.Log `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	// data 数据
	Data *v1.Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// server 服务
	Server *v1.Server `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_config_v1_config_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_api_config_v1_config_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_api_config_v1_config_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetApp() *v1.App {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetLog() *v1.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Bootstrap) GetData() *v1.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetServer() *v1.Server {
	if x != nil {
		return x.Server
	}
	return nil
}

var File_api_config_v1_config_v1_proto protoreflect.FileDescriptor

var file_api_config_v1_config_v1_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x2a,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x69,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x74,
	0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x42, 0x15, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76,
	0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_config_v1_config_v1_proto_rawDescOnce sync.Once
	file_api_config_v1_config_v1_proto_rawDescData = file_api_config_v1_config_v1_proto_rawDesc
)

func file_api_config_v1_config_v1_proto_rawDescGZIP() []byte {
	file_api_config_v1_config_v1_proto_rawDescOnce.Do(func() {
		file_api_config_v1_config_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_config_v1_config_v1_proto_rawDescData)
	})
	return file_api_config_v1_config_v1_proto_rawDescData
}

var file_api_config_v1_config_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_config_v1_config_v1_proto_goTypes = []interface{}{
	(*Bootstrap)(nil), // 0: user.api.config.configv1.Bootstrap
	(*v1.App)(nil),    // 1: kit.api.conf.confv1.App
	(*v1.Log)(nil),    // 2: kit.api.conf.confv1.Log
	(*v1.Data)(nil),   // 3: kit.api.conf.confv1.Data
	(*v1.Server)(nil), // 4: kit.api.conf.confv1.Server
}
var file_api_config_v1_config_v1_proto_depIdxs = []int32{
	1, // 0: user.api.config.configv1.Bootstrap.app:type_name -> kit.api.conf.confv1.App
	2, // 1: user.api.config.configv1.Bootstrap.log:type_name -> kit.api.conf.confv1.Log
	3, // 2: user.api.config.configv1.Bootstrap.data:type_name -> kit.api.conf.confv1.Data
	4, // 3: user.api.config.configv1.Bootstrap.server:type_name -> kit.api.conf.confv1.Server
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_config_v1_config_v1_proto_init() }
func file_api_config_v1_config_v1_proto_init() {
	if File_api_config_v1_config_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_config_v1_config_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
			RawDescriptor: file_api_config_v1_config_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_config_v1_config_v1_proto_goTypes,
		DependencyIndexes: file_api_config_v1_config_v1_proto_depIdxs,
		MessageInfos:      file_api_config_v1_config_v1_proto_msgTypes,
	}.Build()
	File_api_config_v1_config_v1_proto = out.File
	file_api_config_v1_config_v1_proto_rawDesc = nil
	file_api_config_v1_config_v1_proto_goTypes = nil
	file_api_config_v1_config_v1_proto_depIdxs = nil
}
