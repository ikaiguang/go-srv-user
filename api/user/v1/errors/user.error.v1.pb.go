// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/user/v1/errors/user.error.v1.proto

package usererrorv1

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

// UserError 用户错误
type UserErrorEnum_UserError int32

const (
	// UNSPECIFIED 初始化
	UserErrorEnum_UNSPECIFIED UserErrorEnum_UserError = 0
	// NOT_EXIST 用户不存在
	UserErrorEnum_NOT_EXIST UserErrorEnum_UserError = 1
	// EXIST 用户已存在
	UserErrorEnum_EXIST UserErrorEnum_UserError = 2
	// NAME_INVALID 用户名不合法
	UserErrorEnum_NAME_INVALID UserErrorEnum_UserError = 3
	// NAME_EXIST 用户名已存在
	UserErrorEnum_NAME_EXIST UserErrorEnum_UserError = 4
	// PASSWORD_INVALID 用户密码不合法
	UserErrorEnum_PASSWORD_INVALID UserErrorEnum_UserError = 5
)

// Enum value maps for UserErrorEnum_UserError.
var (
	UserErrorEnum_UserError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "NOT_EXIST",
		2: "EXIST",
		3: "NAME_INVALID",
		4: "NAME_EXIST",
		5: "PASSWORD_INVALID",
	}
	UserErrorEnum_UserError_value = map[string]int32{
		"UNSPECIFIED":      0,
		"NOT_EXIST":        1,
		"EXIST":            2,
		"NAME_INVALID":     3,
		"NAME_EXIST":       4,
		"PASSWORD_INVALID": 5,
	}
)

func (x UserErrorEnum_UserError) Enum() *UserErrorEnum_UserError {
	p := new(UserErrorEnum_UserError)
	*p = x
	return p
}

func (x UserErrorEnum_UserError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserErrorEnum_UserError) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_v1_errors_user_error_v1_proto_enumTypes[0].Descriptor()
}

func (UserErrorEnum_UserError) Type() protoreflect.EnumType {
	return &file_api_user_v1_errors_user_error_v1_proto_enumTypes[0]
}

func (x UserErrorEnum_UserError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserErrorEnum_UserError.Descriptor instead.
func (UserErrorEnum_UserError) EnumDescriptor() ([]byte, []int) {
	return file_api_user_v1_errors_user_error_v1_proto_rawDescGZIP(), []int{0, 0}
}

// UserErrorEnum 用户错误
type UserErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserErrorEnum) Reset() {
	*x = UserErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_errors_user_error_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserErrorEnum) ProtoMessage() {}

func (x *UserErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_errors_user_error_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserErrorEnum.ProtoReflect.Descriptor instead.
func (*UserErrorEnum) Descriptor() ([]byte, []int) {
	return file_api_user_v1_errors_user_error_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_user_v1_errors_user_error_v1_proto protoreflect.FileDescriptor

var file_api_user_v1_errors_user_error_v1_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x76, 0x31, 0x22, 0x7f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x6e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x04,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x05, 0x42, 0x68, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x72, 0x76, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x07, 0x55, 0x73, 0x65, 0x72, 0x41, 0x50, 0x49,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_errors_user_error_v1_proto_rawDescOnce sync.Once
	file_api_user_v1_errors_user_error_v1_proto_rawDescData = file_api_user_v1_errors_user_error_v1_proto_rawDesc
)

func file_api_user_v1_errors_user_error_v1_proto_rawDescGZIP() []byte {
	file_api_user_v1_errors_user_error_v1_proto_rawDescOnce.Do(func() {
		file_api_user_v1_errors_user_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_errors_user_error_v1_proto_rawDescData)
	})
	return file_api_user_v1_errors_user_error_v1_proto_rawDescData
}

var file_api_user_v1_errors_user_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_user_v1_errors_user_error_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_user_v1_errors_user_error_v1_proto_goTypes = []interface{}{
	(UserErrorEnum_UserError)(0), // 0: user.api.config.usererrorv1.UserErrorEnum.UserError
	(*UserErrorEnum)(nil),        // 1: user.api.config.usererrorv1.UserErrorEnum
}
var file_api_user_v1_errors_user_error_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_errors_user_error_v1_proto_init() }
func file_api_user_v1_errors_user_error_v1_proto_init() {
	if File_api_user_v1_errors_user_error_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_v1_errors_user_error_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserErrorEnum); i {
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
			RawDescriptor: file_api_user_v1_errors_user_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_v1_errors_user_error_v1_proto_goTypes,
		DependencyIndexes: file_api_user_v1_errors_user_error_v1_proto_depIdxs,
		EnumInfos:         file_api_user_v1_errors_user_error_v1_proto_enumTypes,
		MessageInfos:      file_api_user_v1_errors_user_error_v1_proto_msgTypes,
	}.Build()
	File_api_user_v1_errors_user_error_v1_proto = out.File
	file_api_user_v1_errors_user_error_v1_proto_rawDesc = nil
	file_api_user_v1_errors_user_error_v1_proto_goTypes = nil
	file_api_user_v1_errors_user_error_v1_proto_depIdxs = nil
}
