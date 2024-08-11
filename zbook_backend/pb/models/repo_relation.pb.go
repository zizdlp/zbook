// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: models/repo_relation.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListRepoReportInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportId      int64                  `protobuf:"varint,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RepoId        int64                  `protobuf:"varint,3,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	ReportContent string                 `protobuf:"bytes,4,opt,name=report_content,json=reportContent,proto3" json:"report_content,omitempty"`
	Processed     bool                   `protobuf:"varint,5,opt,name=processed,proto3" json:"processed,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ListRepoReportInfo) Reset() {
	*x = ListRepoReportInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_repo_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepoReportInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepoReportInfo) ProtoMessage() {}

func (x *ListRepoReportInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_repo_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepoReportInfo.ProtoReflect.Descriptor instead.
func (*ListRepoReportInfo) Descriptor() ([]byte, []int) {
	return file_models_repo_relation_proto_rawDescGZIP(), []int{0}
}

func (x *ListRepoReportInfo) GetReportId() int64 {
	if x != nil {
		return x.ReportId
	}
	return 0
}

func (x *ListRepoReportInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListRepoReportInfo) GetRepoId() int64 {
	if x != nil {
		return x.RepoId
	}
	return 0
}

func (x *ListRepoReportInfo) GetReportContent() string {
	if x != nil {
		return x.ReportContent
	}
	return ""
}

func (x *ListRepoReportInfo) GetProcessed() bool {
	if x != nil {
		return x.Processed
	}
	return false
}

func (x *ListRepoReportInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ListUserRepoVisiblityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	IsFollowing   bool                   `protobuf:"varint,3,opt,name=is_following,json=isFollowing,proto3" json:"is_following,omitempty"`
	IsRepoVisible bool                   `protobuf:"varint,4,opt,name=is_repo_visible,json=isRepoVisible,proto3" json:"is_repo_visible,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ListUserRepoVisiblityInfo) Reset() {
	*x = ListUserRepoVisiblityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_repo_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRepoVisiblityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRepoVisiblityInfo) ProtoMessage() {}

func (x *ListUserRepoVisiblityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_repo_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRepoVisiblityInfo.ProtoReflect.Descriptor instead.
func (*ListUserRepoVisiblityInfo) Descriptor() ([]byte, []int) {
	return file_models_repo_relation_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserRepoVisiblityInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListUserRepoVisiblityInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ListUserRepoVisiblityInfo) GetIsFollowing() bool {
	if x != nil {
		return x.IsFollowing
	}
	return false
}

func (x *ListUserRepoVisiblityInfo) GetIsRepoVisible() bool {
	if x != nil {
		return x.IsRepoVisible
	}
	return false
}

func (x *ListUserRepoVisiblityInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_models_repo_relation_proto protoreflect.FileDescriptor

var file_models_repo_relation_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a, 0x64,
	0x6c, 0x70, 0x2f, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_repo_relation_proto_rawDescOnce sync.Once
	file_models_repo_relation_proto_rawDescData = file_models_repo_relation_proto_rawDesc
)

func file_models_repo_relation_proto_rawDescGZIP() []byte {
	file_models_repo_relation_proto_rawDescOnce.Do(func() {
		file_models_repo_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_repo_relation_proto_rawDescData)
	})
	return file_models_repo_relation_proto_rawDescData
}

var file_models_repo_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_repo_relation_proto_goTypes = []interface{}{
	(*ListRepoReportInfo)(nil),        // 0: pb.ListRepoReportInfo
	(*ListUserRepoVisiblityInfo)(nil), // 1: pb.ListUserRepoVisiblityInfo
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
}
var file_models_repo_relation_proto_depIdxs = []int32{
	2, // 0: pb.ListRepoReportInfo.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: pb.ListUserRepoVisiblityInfo.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_repo_relation_proto_init() }
func file_models_repo_relation_proto_init() {
	if File_models_repo_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_repo_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepoReportInfo); i {
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
		file_models_repo_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRepoVisiblityInfo); i {
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
			RawDescriptor: file_models_repo_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_repo_relation_proto_goTypes,
		DependencyIndexes: file_models_repo_relation_proto_depIdxs,
		MessageInfos:      file_models_repo_relation_proto_msgTypes,
	}.Build()
	File_models_repo_relation_proto = out.File
	file_models_repo_relation_proto_rawDesc = nil
	file_models_repo_relation_proto_goTypes = nil
	file_models_repo_relation_proto_depIdxs = nil
}
