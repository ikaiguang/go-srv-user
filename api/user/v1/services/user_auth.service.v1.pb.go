// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/user/v1/services/user_auth.service.v1.proto

package userservicev1

import (
	resources "github.com/ikaiguang/go-srv-user/api/user/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_user_v1_services_user_auth_service_v1_proto protoreflect.FileDescriptor

var file_api_user_v1_services_user_auth_service_v1_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x52, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61,
	0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x9c, 0x01, 0x0a, 0x0b, 0x53, 0x72, 0x76, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x8c, 0x01, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4f, 0x72, 0x52, 0x65, 0x67, 0x42,
	0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4f, 0x72, 0x52, 0x65, 0x67, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2d,
	0x6f, 0x72, 0x2d, 0x72, 0x65, 0x67, 0x2d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x3a, 0x01, 0x2a, 0x42,
	0x7c, 0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x42, 0x16,
	0x54, 0x62, 0x41, 0x70, 0x69, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_user_v1_services_user_auth_service_v1_proto_goTypes = []interface{}{
	(*resources.LoginOrRegByPhoneReq)(nil), // 0: user.api.user.userv1.LoginOrRegByPhoneReq
	(*resources.LoginResp)(nil),            // 1: user.api.user.userv1.LoginResp
}
var file_api_user_v1_services_user_auth_service_v1_proto_depIdxs = []int32{
	0, // 0: user.api.user.userservicev1.SrvUserAuth.LoginOrRegByPhone:input_type -> user.api.user.userv1.LoginOrRegByPhoneReq
	1, // 1: user.api.user.userservicev1.SrvUserAuth.LoginOrRegByPhone:output_type -> user.api.user.userv1.LoginResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_services_user_auth_service_v1_proto_init() }
func file_api_user_v1_services_user_auth_service_v1_proto_init() {
	if File_api_user_v1_services_user_auth_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_v1_services_user_auth_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_v1_services_user_auth_service_v1_proto_goTypes,
		DependencyIndexes: file_api_user_v1_services_user_auth_service_v1_proto_depIdxs,
	}.Build()
	File_api_user_v1_services_user_auth_service_v1_proto = out.File
	file_api_user_v1_services_user_auth_service_v1_proto_rawDesc = nil
	file_api_user_v1_services_user_auth_service_v1_proto_goTypes = nil
	file_api_user_v1_services_user_auth_service_v1_proto_depIdxs = nil
}