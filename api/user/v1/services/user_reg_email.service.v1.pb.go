// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/user/v1/services/user_reg_email.service.v1.proto

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

var File_api_user_v1_services_user_reg_email_service_v1_proto protoreflect.FileDescriptor

var file_api_user_v1_services_user_reg_email_service_v1_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b,
	0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x67, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xab, 0x06, 0x0a, 0x0f, 0x53,
	0x72, 0x76, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x7c,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x27, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x80, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x2a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x7f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x7f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x1a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x87, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x07,
	0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x2f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x2f, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x42, 0x7c, 0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x42, 0x16, 0x54, 0x62, 0x41, 0x70, 0x69, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b,
	0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_user_v1_services_user_reg_email_service_v1_proto_goTypes = []interface{}{
	(*resources.UserRegEmailIdReq)(nil),         // 0: user.api.user.userv1.UserRegEmailIdReq
	(*resources.UserRegEmailListReq)(nil),       // 1: user.api.user.userv1.UserRegEmailListReq
	(*resources.UserRegEmailSaveReq)(nil),       // 2: user.api.user.userv1.UserRegEmailSaveReq
	(*resources.UserRegEmailIdsReq)(nil),        // 3: user.api.user.userv1.UserRegEmailIdsReq
	(*resources.UserRegEmail)(nil),              // 4: user.api.user.userv1.UserRegEmail
	(*resources.UserRegEmailListResp)(nil),      // 5: user.api.user.userv1.UserRegEmailListResp
	(*resources.UserRegEmailProcessResult)(nil), // 6: user.api.user.userv1.UserRegEmailProcessResult
}
var file_api_user_v1_services_user_reg_email_service_v1_proto_depIdxs = []int32{
	0, // 0: user.api.user.userservicev1.SrvUserRegEmail.GetById:input_type -> user.api.user.userv1.UserRegEmailIdReq
	1, // 1: user.api.user.userservicev1.SrvUserRegEmail.List:input_type -> user.api.user.userv1.UserRegEmailListReq
	2, // 2: user.api.user.userservicev1.SrvUserRegEmail.Create:input_type -> user.api.user.userv1.UserRegEmailSaveReq
	2, // 3: user.api.user.userservicev1.SrvUserRegEmail.Update:input_type -> user.api.user.userv1.UserRegEmailSaveReq
	0, // 4: user.api.user.userservicev1.SrvUserRegEmail.Delete:input_type -> user.api.user.userv1.UserRegEmailIdReq
	3, // 5: user.api.user.userservicev1.SrvUserRegEmail.Destroy:input_type -> user.api.user.userv1.UserRegEmailIdsReq
	4, // 6: user.api.user.userservicev1.SrvUserRegEmail.GetById:output_type -> user.api.user.userv1.UserRegEmail
	5, // 7: user.api.user.userservicev1.SrvUserRegEmail.List:output_type -> user.api.user.userv1.UserRegEmailListResp
	4, // 8: user.api.user.userservicev1.SrvUserRegEmail.Create:output_type -> user.api.user.userv1.UserRegEmail
	4, // 9: user.api.user.userservicev1.SrvUserRegEmail.Update:output_type -> user.api.user.userv1.UserRegEmail
	6, // 10: user.api.user.userservicev1.SrvUserRegEmail.Delete:output_type -> user.api.user.userv1.UserRegEmailProcessResult
	6, // 11: user.api.user.userservicev1.SrvUserRegEmail.Destroy:output_type -> user.api.user.userv1.UserRegEmailProcessResult
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_services_user_reg_email_service_v1_proto_init() }
func file_api_user_v1_services_user_reg_email_service_v1_proto_init() {
	if File_api_user_v1_services_user_reg_email_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_v1_services_user_reg_email_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_v1_services_user_reg_email_service_v1_proto_goTypes,
		DependencyIndexes: file_api_user_v1_services_user_reg_email_service_v1_proto_depIdxs,
	}.Build()
	File_api_user_v1_services_user_reg_email_service_v1_proto = out.File
	file_api_user_v1_services_user_reg_email_service_v1_proto_rawDesc = nil
	file_api_user_v1_services_user_reg_email_service_v1_proto_goTypes = nil
	file_api_user_v1_services_user_reg_email_service_v1_proto_depIdxs = nil
}