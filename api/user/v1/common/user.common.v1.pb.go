// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/user/v1/common/user.common.v1.proto

package usercommonv1

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

type AuthTypeEnum_AuthType int32

const (
	// USER_PASSWORD 用户名密码
	AuthTypeEnum_USER_PASSWORD AuthTypeEnum_AuthType = 0
	// SMS_CODE 短信验证码
	AuthTypeEnum_SMS_CODE AuthTypeEnum_AuthType = 1
	// WECHAT_LOGIN 微信登录
	AuthTypeEnum_WECHAT_LOGIN AuthTypeEnum_AuthType = 2
	// ALIPAY_LOGIN 支付宝登录
	AuthTypeEnum_ALIPAY_LOGIN AuthTypeEnum_AuthType = 3
	// TAOBAO_LOGIN 淘宝登录
	AuthTypeEnum_TAOBAO_LOGIN AuthTypeEnum_AuthType = 4
)

// Enum value maps for AuthTypeEnum_AuthType.
var (
	AuthTypeEnum_AuthType_name = map[int32]string{
		0: "USER_PASSWORD",
		1: "SMS_CODE",
		2: "WECHAT_LOGIN",
		3: "ALIPAY_LOGIN",
		4: "TAOBAO_LOGIN",
	}
	AuthTypeEnum_AuthType_value = map[string]int32{
		"USER_PASSWORD": 0,
		"SMS_CODE":      1,
		"WECHAT_LOGIN":  2,
		"ALIPAY_LOGIN":  3,
		"TAOBAO_LOGIN":  4,
	}
)

func (x AuthTypeEnum_AuthType) Enum() *AuthTypeEnum_AuthType {
	p := new(AuthTypeEnum_AuthType)
	*p = x
	return p
}

func (x AuthTypeEnum_AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthTypeEnum_AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_v1_common_user_common_v1_proto_enumTypes[0].Descriptor()
}

func (AuthTypeEnum_AuthType) Type() protoreflect.EnumType {
	return &file_api_user_v1_common_user_common_v1_proto_enumTypes[0]
}

func (x AuthTypeEnum_AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthTypeEnum_AuthType.Descriptor instead.
func (AuthTypeEnum_AuthType) EnumDescriptor() ([]byte, []int) {
	return file_api_user_v1_common_user_common_v1_proto_rawDescGZIP(), []int{0, 0}
}

// AuthTypeEnum 认证类型
type AuthTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthTypeEnum) Reset() {
	*x = AuthTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_common_user_common_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTypeEnum) ProtoMessage() {}

func (x *AuthTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_common_user_common_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTypeEnum.ProtoReflect.Descriptor instead.
func (*AuthTypeEnum) Descriptor() ([]byte, []int) {
	return file_api_user_v1_common_user_common_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_user_v1_common_user_common_v1_proto protoreflect.FileDescriptor

var file_api_user_v1_common_user_common_v1_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x22, 0x71, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x61, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f,
	0x52, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4d, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4c, 0x49, 0x50, 0x41, 0x59, 0x5f, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x4f, 0x42, 0x41, 0x4f,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x04, 0x42, 0x79, 0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x42, 0x17, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x56, 0x31, 0x50,
	0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b,
	0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_common_user_common_v1_proto_rawDescOnce sync.Once
	file_api_user_v1_common_user_common_v1_proto_rawDescData = file_api_user_v1_common_user_common_v1_proto_rawDesc
)

func file_api_user_v1_common_user_common_v1_proto_rawDescGZIP() []byte {
	file_api_user_v1_common_user_common_v1_proto_rawDescOnce.Do(func() {
		file_api_user_v1_common_user_common_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_common_user_common_v1_proto_rawDescData)
	})
	return file_api_user_v1_common_user_common_v1_proto_rawDescData
}

var file_api_user_v1_common_user_common_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_user_v1_common_user_common_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_user_v1_common_user_common_v1_proto_goTypes = []interface{}{
	(AuthTypeEnum_AuthType)(0), // 0: user.api.user.usercommonv1.AuthTypeEnum.AuthType
	(*AuthTypeEnum)(nil),       // 1: user.api.user.usercommonv1.AuthTypeEnum
}
var file_api_user_v1_common_user_common_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_common_user_common_v1_proto_init() }
func file_api_user_v1_common_user_common_v1_proto_init() {
	if File_api_user_v1_common_user_common_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_v1_common_user_common_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthTypeEnum); i {
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
			RawDescriptor: file_api_user_v1_common_user_common_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_v1_common_user_common_v1_proto_goTypes,
		DependencyIndexes: file_api_user_v1_common_user_common_v1_proto_depIdxs,
		EnumInfos:         file_api_user_v1_common_user_common_v1_proto_enumTypes,
		MessageInfos:      file_api_user_v1_common_user_common_v1_proto_msgTypes,
	}.Build()
	File_api_user_v1_common_user_common_v1_proto = out.File
	file_api_user_v1_common_user_common_v1_proto_rawDesc = nil
	file_api_user_v1_common_user_common_v1_proto_goTypes = nil
	file_api_user_v1_common_user_common_v1_proto_depIdxs = nil
}
