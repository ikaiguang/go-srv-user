// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/org/v1/services/org.service.v1.proto

package orgservicev1

import (
	resources "github.com/ikaiguang/go-srv-user/api/org/v1/resources"
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

var File_api_org_v1_services_org_service_v1_proto protoreflect.FileDescriptor

var file_api_org_v1_services_org_service_v1_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76,
	0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xe5, 0x04, 0x0a, 0x06, 0x53, 0x72, 0x76, 0x4f, 0x72, 0x67, 0x12, 0x5d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x4f,
	0x72, 0x67, 0x2f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x61, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x4f, 0x72, 0x67, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x60, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x4f, 0x72, 0x67, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x60, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x67, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x67, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x1a, 0x12, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x4f, 0x72, 0x67, 0x2f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x68, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x4f, 0x72, 0x67, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x6b, 0x0a, 0x07, 0x44,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x49,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x4f, 0x72, 0x67,
	0x2f, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x42, 0x76, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0x42, 0x14, 0x54, 0x62, 0x41, 0x70, 0x69, 0x4f, 0x72, 0x67, 0x4f,
	0x72, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75,
	0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3b, 0x6f, 0x72, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_org_v1_services_org_service_v1_proto_goTypes = []interface{}{
	(*resources.OrgIdReq)(nil),         // 0: user.api.org.orgv1.OrgIdReq
	(*resources.OrgListReq)(nil),       // 1: user.api.org.orgv1.OrgListReq
	(*resources.OrgSaveReq)(nil),       // 2: user.api.org.orgv1.OrgSaveReq
	(*resources.OrgIdsReq)(nil),        // 3: user.api.org.orgv1.OrgIdsReq
	(*resources.Org)(nil),              // 4: user.api.org.orgv1.Org
	(*resources.OrgListResp)(nil),      // 5: user.api.org.orgv1.OrgListResp
	(*resources.OrgProcessResult)(nil), // 6: user.api.org.orgv1.OrgProcessResult
}
var file_api_org_v1_services_org_service_v1_proto_depIdxs = []int32{
	0, // 0: user.api.org.orgservicev1.SrvOrg.GetById:input_type -> user.api.org.orgv1.OrgIdReq
	1, // 1: user.api.org.orgservicev1.SrvOrg.List:input_type -> user.api.org.orgv1.OrgListReq
	2, // 2: user.api.org.orgservicev1.SrvOrg.Create:input_type -> user.api.org.orgv1.OrgSaveReq
	2, // 3: user.api.org.orgservicev1.SrvOrg.Update:input_type -> user.api.org.orgv1.OrgSaveReq
	0, // 4: user.api.org.orgservicev1.SrvOrg.Delete:input_type -> user.api.org.orgv1.OrgIdReq
	3, // 5: user.api.org.orgservicev1.SrvOrg.Destroy:input_type -> user.api.org.orgv1.OrgIdsReq
	4, // 6: user.api.org.orgservicev1.SrvOrg.GetById:output_type -> user.api.org.orgv1.Org
	5, // 7: user.api.org.orgservicev1.SrvOrg.List:output_type -> user.api.org.orgv1.OrgListResp
	4, // 8: user.api.org.orgservicev1.SrvOrg.Create:output_type -> user.api.org.orgv1.Org
	4, // 9: user.api.org.orgservicev1.SrvOrg.Update:output_type -> user.api.org.orgv1.Org
	6, // 10: user.api.org.orgservicev1.SrvOrg.Delete:output_type -> user.api.org.orgv1.OrgProcessResult
	6, // 11: user.api.org.orgservicev1.SrvOrg.Destroy:output_type -> user.api.org.orgv1.OrgProcessResult
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_org_v1_services_org_service_v1_proto_init() }
func file_api_org_v1_services_org_service_v1_proto_init() {
	if File_api_org_v1_services_org_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_org_v1_services_org_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_org_v1_services_org_service_v1_proto_goTypes,
		DependencyIndexes: file_api_org_v1_services_org_service_v1_proto_depIdxs,
	}.Build()
	File_api_org_v1_services_org_service_v1_proto = out.File
	file_api_org_v1_services_org_service_v1_proto_rawDesc = nil
	file_api_org_v1_services_org_service_v1_proto_goTypes = nil
	file_api_org_v1_services_org_service_v1_proto_depIdxs = nil
}
