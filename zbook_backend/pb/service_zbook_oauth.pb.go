// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.3
// source: service_zbook_oauth.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	rpcs "github.com/zizdlp/zbook/pb/rpcs"
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

var File_service_zbook_oauth_proto protoreflect.FileDescriptor

var file_service_zbook_oauth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72,
	0x70, 0x63, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x05, 0x0a, 0x0a, 0x5a, 0x42, 0x6f, 0x6f, 0x6b, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x12, 0xae, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x62, 0x92, 0x41, 0x3f, 0x12, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77,
	0x20, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x26, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x20, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0xab, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5c, 0x92, 0x41, 0x38, 0x12, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x22, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x20, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0xa6, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a,
	0x92, 0x41, 0x37, 0x12, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x22, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x98, 0x01, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42,
	0x79, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55,
	0x92, 0x41, 0x35, 0x12, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x20, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01,
	0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x5f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x42, 0x6d, 0x92, 0x41, 0x4e, 0x12, 0x4c, 0x0a, 0x09, 0x7a, 0x62,
	0x6f, 0x6f, 0x6b, 0x20, 0x61, 0x70, 0x69, 0x22, 0x3a, 0x0a, 0x05, 0x7a, 0x62, 0x6f, 0x6f, 0x6b,
	0x12, 0x1f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x2f, 0x7a, 0x62, 0x6f, 0x6f,
	0x6b, 0x1a, 0x10, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x63, 0x6f, 0x6d, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x2f, 0x7a, 0x62, 0x6f, 0x6f,
	0x6b, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_zbook_oauth_proto_goTypes = []interface{}{
	(*rpcs.CreateOAuthLinkRequest)(nil),   // 0: pb.CreateOAuthLinkRequest
	(*rpcs.CheckOAuthStatusRequest)(nil),  // 1: pb.CheckOAuthStatusRequest
	(*rpcs.DeleteOAuthLinkRequest)(nil),   // 2: pb.DeleteOAuthLinkRequest
	(*rpcs.LoginByOAuthRequest)(nil),      // 3: pb.LoginByOAuthRequest
	(*rpcs.CreateOAuthLinkResponse)(nil),  // 4: pb.CreateOAuthLinkResponse
	(*rpcs.CheckOAuthStatusResponse)(nil), // 5: pb.CheckOAuthStatusResponse
	(*rpcs.DeleteOAuthLinkResponse)(nil),  // 6: pb.DeleteOAuthLinkResponse
	(*rpcs.LoginByOAuthResponse)(nil),     // 7: pb.LoginByOAuthResponse
}
var file_service_zbook_oauth_proto_depIdxs = []int32{
	0, // 0: pb.ZBookOAuth.CreateOAuthLink:input_type -> pb.CreateOAuthLinkRequest
	1, // 1: pb.ZBookOAuth.CheckOAuthStatus:input_type -> pb.CheckOAuthStatusRequest
	2, // 2: pb.ZBookOAuth.DeleteOAuthLink:input_type -> pb.DeleteOAuthLinkRequest
	3, // 3: pb.ZBookOAuth.LoginByOAuth:input_type -> pb.LoginByOAuthRequest
	4, // 4: pb.ZBookOAuth.CreateOAuthLink:output_type -> pb.CreateOAuthLinkResponse
	5, // 5: pb.ZBookOAuth.CheckOAuthStatus:output_type -> pb.CheckOAuthStatusResponse
	6, // 6: pb.ZBookOAuth.DeleteOAuthLink:output_type -> pb.DeleteOAuthLinkResponse
	7, // 7: pb.ZBookOAuth.LoginByOAuth:output_type -> pb.LoginByOAuthResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_zbook_oauth_proto_init() }
func file_service_zbook_oauth_proto_init() {
	if File_service_zbook_oauth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_zbook_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_zbook_oauth_proto_goTypes,
		DependencyIndexes: file_service_zbook_oauth_proto_depIdxs,
	}.Build()
	File_service_zbook_oauth_proto = out.File
	file_service_zbook_oauth_proto_rawDesc = nil
	file_service_zbook_oauth_proto_goTypes = nil
	file_service_zbook_oauth_proto_depIdxs = nil
}
