// clang-format off

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: service_zbook_verification.proto

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

var File_service_zbook_verification_proto protoreflect.FileDescriptor

var file_service_zbook_verification_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xc4, 0x07, 0x0a, 0x11, 0x5a, 0x42, 0x6f, 0x6f, 0x6b, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x96, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x3b, 0x12, 0x0c, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2b, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x96, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x92, 0x41, 0x30, 0x12, 0x0e, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x1e, 0x55, 0x73,
	0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0xf6, 0x01, 0x0a, 0x18, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x8e, 0x01, 0x92, 0x41, 0x60, 0x12, 0x2e, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x6e, 0x64, 0x20, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x79, 0x6f, 0x75,
	0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2e, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x6e, 0x64, 0x20, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x79, 0x6f, 0x75,
	0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a,
	0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0xf0, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x54, 0x6f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54,
	0x6f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8e, 0x01, 0x92, 0x41, 0x62, 0x12, 0x2f, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x6e, 0x64,
	0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x74, 0x6f,
	0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2f, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x6e,
	0x64, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x74,
	0x6f, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x90, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x92, 0x41, 0x2e, 0x12, 0x0d,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1d, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6d, 0x92, 0x41, 0x4e, 0x12, 0x4c, 0x0a,
	0x09, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x20, 0x61, 0x70, 0x69, 0x22, 0x3a, 0x0a, 0x05, 0x7a, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x2f, 0x7a,
	0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x10, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x40, 0x67, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x5a, 0x1a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x2f, 0x7a,
	0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_zbook_verification_proto_goTypes = []interface{}{
	(*rpcs.VerifyEmailRequest)(nil),               // 0: pb.VerifyEmailRequest
	(*rpcs.ResetPasswordRequest)(nil),             // 1: pb.ResetPasswordRequest
	(*rpcs.SendEmailToResetPasswordRequest)(nil),  // 2: pb.SendEmailToResetPasswordRequest
	(*rpcs.SendEmailToVerifyEmailRequest)(nil),    // 3: pb.SendEmailToVerifyEmailRequest
	(*rpcs.RefreshTokenRequest)(nil),              // 4: pb.RefreshTokenRequest
	(*rpcs.VerifyEmailResponse)(nil),              // 5: pb.VerifyEmailResponse
	(*rpcs.ResetPasswordResponse)(nil),            // 6: pb.ResetPasswordResponse
	(*rpcs.SendEmailToResetPasswordResponse)(nil), // 7: pb.SendEmailToResetPasswordResponse
	(*rpcs.SendEmailToVerifyEmailResponse)(nil),   // 8: pb.SendEmailToVerifyEmailResponse
	(*rpcs.RefreshTokenResponse)(nil),             // 9: pb.RefreshTokenResponse
}
var file_service_zbook_verification_proto_depIdxs = []int32{
	0, // 0: pb.ZBookVerification.VerifyEmail:input_type -> pb.VerifyEmailRequest
	1, // 1: pb.ZBookVerification.ResetPassword:input_type -> pb.ResetPasswordRequest
	2, // 2: pb.ZBookVerification.SendEmailToResetPassword:input_type -> pb.SendEmailToResetPasswordRequest
	3, // 3: pb.ZBookVerification.SendEmailToVerifyEmail:input_type -> pb.SendEmailToVerifyEmailRequest
	4, // 4: pb.ZBookVerification.RefreshToken:input_type -> pb.RefreshTokenRequest
	5, // 5: pb.ZBookVerification.VerifyEmail:output_type -> pb.VerifyEmailResponse
	6, // 6: pb.ZBookVerification.ResetPassword:output_type -> pb.ResetPasswordResponse
	7, // 7: pb.ZBookVerification.SendEmailToResetPassword:output_type -> pb.SendEmailToResetPasswordResponse
	8, // 8: pb.ZBookVerification.SendEmailToVerifyEmail:output_type -> pb.SendEmailToVerifyEmailResponse
	9, // 9: pb.ZBookVerification.RefreshToken:output_type -> pb.RefreshTokenResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_zbook_verification_proto_init() }
func file_service_zbook_verification_proto_init() {
	if File_service_zbook_verification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_zbook_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_zbook_verification_proto_goTypes,
		DependencyIndexes: file_service_zbook_verification_proto_depIdxs,
	}.Build()
	File_service_zbook_verification_proto = out.File
	file_service_zbook_verification_proto_rawDesc = nil
	file_service_zbook_verification_proto_goTypes = nil
	file_service_zbook_verification_proto_depIdxs = nil
}
