// clang-format off

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: service_zbook_notification.proto

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

var File_service_zbook_notification_proto protoreflect.FileDescriptor

var file_service_zbook_notification_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xf4, 0x1a, 0x0a, 0x11, 0x5a, 0x42, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x84, 0x02, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x9c, 0x01, 0x92, 0x41, 0x70, 0x12, 0x36, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97,
	0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7,
	0x9a, 0x84, 0xe5, 0x85, 0xb3, 0xe6, 0xb3, 0xa8, 0xe8, 0x80, 0x85, 0xe9, 0x80, 0x9a, 0xe7, 0x9f,
	0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x1a, 0x36, 0xe5,
	0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae,
	0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe5, 0x85, 0xb3, 0xe6, 0xb3, 0xa8,
	0xe8, 0x80, 0x85, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4,
	0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xee,
	0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x92, 0x01, 0x92, 0x41, 0x6a,
	0x12, 0x33, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6, 0x8c,
	0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe5, 0xb8, 0x96,
	0xe5, 0xad, 0x90, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4,
	0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x1a, 0x33, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97,
	0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7,
	0x9a, 0x84, 0xe5, 0xb8, 0x96, 0xe5, 0xad, 0x90, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd,
	0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0xfa, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95, 0x01, 0x92, 0x41, 0x6a, 0x12, 0x33, 0xe5, 0x88, 0x86, 0xe9,
	0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94,
	0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe8, 0xaf, 0x84, 0xe8, 0xae, 0xba, 0xe9, 0x80, 0x9a,
	0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x1a,
	0x33, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87,
	0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe8, 0xaf, 0x84, 0xe8,
	0xae, 0xba, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf,
	0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xdd, 0x01, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7c,
	0x92, 0x41, 0x52, 0x12, 0x27, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f,
	0x96, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84,
	0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0x1a, 0x27, 0xe5, 0x88,
	0x86, 0xe9, 0xa1, 0xb5, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a,
	0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe9,
	0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xe4, 0x01, 0x0a,
	0x1e, 0x4d, 0x61, 0x72, 0x6b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x12,
	0x29, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x7e, 0x92, 0x41, 0x4b, 0x12, 0x21, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x1a, 0x26, 0x55, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x74,
	0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61,
	0x64, 0x65, 0x64, 0x12, 0xcf, 0x01, 0x0a, 0x1c, 0x4d, 0x61, 0x72, 0x6b, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x64, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x92, 0x41, 0x3c, 0x12, 0x14, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x61,
	0x64, 0x1a, 0x24, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20,
	0x74, 0x6f, 0x20, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a,
	0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x64, 0x12, 0xd4, 0x01, 0x0a, 0x1d, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x70, 0x92, 0x41, 0x3e, 0x12,
	0x15, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x1a, 0x25, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x12, 0xc7, 0x01, 0x0a,
	0x1a, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41, 0x3a,
	0x12, 0x13, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x1a, 0x23, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26,
	0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x12, 0x9f, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x6e,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x52,
	0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x56, 0x92, 0x41, 0x34, 0x12, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x6e, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a,
	0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x6e, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0xab, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x92, 0x41, 0x38, 0x12, 0x12, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x1a, 0x22, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20,
	0x74, 0x6f, 0x20, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0xc7, 0x02, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xaf,
	0x01, 0x92, 0x41, 0x70, 0x12, 0x36, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5,
	0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a,
	0x84, 0xe5, 0x85, 0xb3, 0xe6, 0xb3, 0xa8, 0xe8, 0x80, 0x85, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5,
	0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x1a, 0x36, 0xe5, 0x88,
	0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a,
	0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe5, 0x85, 0xb3, 0xe6, 0xb3, 0xa8, 0xe8,
	0x80, 0x85, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf,
	0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x3a, 0x01, 0x2a, 0x22, 0x31, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0xb1, 0x02, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa5, 0x01, 0x92,
	0x41, 0x6a, 0x12, 0x33, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba,
	0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe5,
	0xb8, 0x96, 0xe5, 0xad, 0x90, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1,
	0x9e, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x1a, 0x33, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5,
	0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88,
	0xb7, 0xe7, 0x9a, 0x84, 0xe5, 0xb8, 0x96, 0xe5, 0xad, 0x90, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5,
	0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x32, 0x3a, 0x01, 0x2a, 0x22, 0x2d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0xbd, 0x02, 0x0a, 0x27, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x32, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa8, 0x01, 0x92, 0x41, 0x6a, 0x12,
	0x33, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5, 0x87, 0xba, 0xe6, 0x8c, 0x87,
	0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe8, 0xaf, 0x84, 0xe8,
	0xae, 0xba, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92, 0xe5, 0xb1, 0x9e, 0xe4, 0xbf,
	0xa1, 0xe6, 0x81, 0xaf, 0x1a, 0x33, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe5, 0x88, 0x97, 0xe5,
	0x87, 0xba, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a,
	0x84, 0xe8, 0xaf, 0x84, 0xe8, 0xae, 0xba, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0xe5, 0xbd, 0x92,
	0xe5, 0xb1, 0x9e, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x3a,
	0x01, 0x2a, 0x22, 0x30, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0xa1, 0x02, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x31, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x92, 0x41, 0x52, 0x12, 0x27, 0xe5, 0x88,
	0x86, 0xe9, 0xa1, 0xb5, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a,
	0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7, 0x9a, 0x84, 0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe9,
	0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0x1a, 0x27, 0xe5, 0x88, 0x86, 0xe9, 0xa1, 0xb5, 0xe8, 0x8e, 0xb7,
	0xe5, 0x8f, 0x96, 0xe6, 0x8c, 0x87, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe7,
	0x9a, 0x84, 0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe9, 0x80, 0x9a, 0xe7, 0x9f, 0xa5, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x34, 0x3a, 0x01, 0x2a, 0x22, 0x2f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x72, 0x92, 0x41, 0x53, 0x12, 0x51, 0x0a,
	0x09, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x20, 0x61, 0x70, 0x69, 0x22, 0x3f, 0x0a, 0x0a, 0x7a, 0x69,
	0x7a, 0x64, 0x6c, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x12, 0x1f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a,
	0x64, 0x6c, 0x70, 0x2f, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x10, 0x7a, 0x69, 0x7a, 0x64, 0x6c,
	0x70, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x30, 0x2e, 0x31,
	0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a,
	0x64, 0x6c, 0x70, 0x2f, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_service_zbook_notification_proto_goTypes = []interface{}{
	(*rpcs.ListFollowerNotificationRequest)(nil),                  // 0: pb.ListFollowerNotificationRequest
	(*rpcs.ListRepoNotificationRequest)(nil),                      // 1: pb.ListRepoNotificationRequest
	(*rpcs.ListCommentNotificationRequest)(nil),                   // 2: pb.ListCommentNotificationRequest
	(*rpcs.ListSystemNotificationRequest)(nil),                    // 3: pb.ListSystemNotificationRequest
	(*rpcs.MarkFollowerNotificationReadedRequest)(nil),            // 4: pb.MarkFollowerNotificationReadedRequest
	(*rpcs.MarkSystemNotificationReadedRequest)(nil),              // 5: pb.MarkSystemNotificationReadedRequest
	(*rpcs.MarkCommentNotificationReadedRequest)(nil),             // 6: pb.MarkCommentNotificationReadedRequest
	(*rpcs.MarkRepoNotificationReadedRequest)(nil),                // 7: pb.MarkRepoNotificationReadedRequest
	(*rpcs.GetUnReadCountRequest)(nil),                            // 8: pb.GetUnReadCountRequest
	(*rpcs.ResetUnreadCountRequest)(nil),                          // 9: pb.ResetUnreadCountRequest
	(*rpcs.GetListFollowerNotificationUnreadedCountRequest)(nil),  // 10: pb.GetListFollowerNotificationUnreadedCountRequest
	(*rpcs.GetListRepoNotificationUnreadedCountRequest)(nil),      // 11: pb.GetListRepoNotificationUnreadedCountRequest
	(*rpcs.GetListCommentNotificationUnreadedCountRequest)(nil),   // 12: pb.GetListCommentNotificationUnreadedCountRequest
	(*rpcs.GetListSystemNotificationUnreadedCountRequest)(nil),    // 13: pb.GetListSystemNotificationUnreadedCountRequest
	(*rpcs.ListFollowerNotificationResponse)(nil),                 // 14: pb.ListFollowerNotificationResponse
	(*rpcs.ListRepoNotificationResponse)(nil),                     // 15: pb.ListRepoNotificationResponse
	(*rpcs.ListCommentNotificationResponse)(nil),                  // 16: pb.ListCommentNotificationResponse
	(*rpcs.ListSystemNotificationResponse)(nil),                   // 17: pb.ListSystemNotificationResponse
	(*rpcs.SetNotiReadResponse)(nil),                              // 18: pb.SetNotiReadResponse
	(*rpcs.GetUnReadCountResponse)(nil),                           // 19: pb.GetUnReadCountResponse
	(*rpcs.ResetUnreadCountResponse)(nil),                         // 20: pb.ResetUnreadCountResponse
	(*rpcs.GetListFollowerNotificationUnreadedCountResponse)(nil), // 21: pb.GetListFollowerNotificationUnreadedCountResponse
	(*rpcs.GetListRepoNotificationUnreadedCountResponse)(nil),     // 22: pb.GetListRepoNotificationUnreadedCountResponse
	(*rpcs.GetListCommentNotificationUnreadedCountResponse)(nil),  // 23: pb.GetListCommentNotificationUnreadedCountResponse
	(*rpcs.GetListSystemNotificationUnreadedCountResponse)(nil),   // 24: pb.GetListSystemNotificationUnreadedCountResponse
}
var file_service_zbook_notification_proto_depIdxs = []int32{
	0,  // 0: pb.ZBookNotification.ListFollowerNotification:input_type -> pb.ListFollowerNotificationRequest
	1,  // 1: pb.ZBookNotification.ListRepoNotification:input_type -> pb.ListRepoNotificationRequest
	2,  // 2: pb.ZBookNotification.ListCommentNotification:input_type -> pb.ListCommentNotificationRequest
	3,  // 3: pb.ZBookNotification.ListSystemNotification:input_type -> pb.ListSystemNotificationRequest
	4,  // 4: pb.ZBookNotification.MarkFollowerNotificationReaded:input_type -> pb.MarkFollowerNotificationReadedRequest
	5,  // 5: pb.ZBookNotification.MarkSystemNotificationReaded:input_type -> pb.MarkSystemNotificationReadedRequest
	6,  // 6: pb.ZBookNotification.MarkCommentNotificationReaded:input_type -> pb.MarkCommentNotificationReadedRequest
	7,  // 7: pb.ZBookNotification.MarkRepoNotificationReaded:input_type -> pb.MarkRepoNotificationReadedRequest
	8,  // 8: pb.ZBookNotification.GetUnReadCount:input_type -> pb.GetUnReadCountRequest
	9,  // 9: pb.ZBookNotification.ResetUnreadCount:input_type -> pb.ResetUnreadCountRequest
	10, // 10: pb.ZBookNotification.GetListFollowerNotificationUnreadedCount:input_type -> pb.GetListFollowerNotificationUnreadedCountRequest
	11, // 11: pb.ZBookNotification.GetListRepoNotificationUnreadedCount:input_type -> pb.GetListRepoNotificationUnreadedCountRequest
	12, // 12: pb.ZBookNotification.GetListCommentNotificationUnreadedCount:input_type -> pb.GetListCommentNotificationUnreadedCountRequest
	13, // 13: pb.ZBookNotification.GetListSystemNotificationUnreadedCount:input_type -> pb.GetListSystemNotificationUnreadedCountRequest
	14, // 14: pb.ZBookNotification.ListFollowerNotification:output_type -> pb.ListFollowerNotificationResponse
	15, // 15: pb.ZBookNotification.ListRepoNotification:output_type -> pb.ListRepoNotificationResponse
	16, // 16: pb.ZBookNotification.ListCommentNotification:output_type -> pb.ListCommentNotificationResponse
	17, // 17: pb.ZBookNotification.ListSystemNotification:output_type -> pb.ListSystemNotificationResponse
	18, // 18: pb.ZBookNotification.MarkFollowerNotificationReaded:output_type -> pb.SetNotiReadResponse
	18, // 19: pb.ZBookNotification.MarkSystemNotificationReaded:output_type -> pb.SetNotiReadResponse
	18, // 20: pb.ZBookNotification.MarkCommentNotificationReaded:output_type -> pb.SetNotiReadResponse
	18, // 21: pb.ZBookNotification.MarkRepoNotificationReaded:output_type -> pb.SetNotiReadResponse
	19, // 22: pb.ZBookNotification.GetUnReadCount:output_type -> pb.GetUnReadCountResponse
	20, // 23: pb.ZBookNotification.ResetUnreadCount:output_type -> pb.ResetUnreadCountResponse
	21, // 24: pb.ZBookNotification.GetListFollowerNotificationUnreadedCount:output_type -> pb.GetListFollowerNotificationUnreadedCountResponse
	22, // 25: pb.ZBookNotification.GetListRepoNotificationUnreadedCount:output_type -> pb.GetListRepoNotificationUnreadedCountResponse
	23, // 26: pb.ZBookNotification.GetListCommentNotificationUnreadedCount:output_type -> pb.GetListCommentNotificationUnreadedCountResponse
	24, // 27: pb.ZBookNotification.GetListSystemNotificationUnreadedCount:output_type -> pb.GetListSystemNotificationUnreadedCountResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_zbook_notification_proto_init() }
func file_service_zbook_notification_proto_init() {
	if File_service_zbook_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_zbook_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_zbook_notification_proto_goTypes,
		DependencyIndexes: file_service_zbook_notification_proto_depIdxs,
	}.Build()
	File_service_zbook_notification_proto = out.File
	file_service_zbook_notification_proto_rawDesc = nil
	file_service_zbook_notification_proto_goTypes = nil
	file_service_zbook_notification_proto_depIdxs = nil
}
