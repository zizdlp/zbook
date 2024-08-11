// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: rpcs/rpc_markdown.proto

package rpcs

import (
	models "github.com/zizdlp/zbook/pb/models"
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

// 1.GetMarkdownContent
type GetMarkdownContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	RepoName     string `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	RelativePath string `protobuf:"bytes,3,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
}

func (x *GetMarkdownContentRequest) Reset() {
	*x = GetMarkdownContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarkdownContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarkdownContentRequest) ProtoMessage() {}

func (x *GetMarkdownContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarkdownContentRequest.ProtoReflect.Descriptor instead.
func (*GetMarkdownContentRequest) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{0}
}

func (x *GetMarkdownContentRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetMarkdownContentRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *GetMarkdownContentRequest) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

type GetMarkdownContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Markdown   *models.Markdown       `protobuf:"bytes,1,opt,name=markdown,proto3" json:"markdown,omitempty"`
	Prev       string                 `protobuf:"bytes,2,opt,name=prev,proto3" json:"prev,omitempty"`
	Next       string                 `protobuf:"bytes,3,opt,name=next,proto3" json:"next,omitempty"`
	Footers    []*models.FooterSocial `protobuf:"bytes,4,rep,name=footers,proto3" json:"footers,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ThemeColor string                 `protobuf:"bytes,6,opt,name=theme_color,json=themeColor,proto3" json:"theme_color,omitempty"`
}

func (x *GetMarkdownContentResponse) Reset() {
	*x = GetMarkdownContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarkdownContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarkdownContentResponse) ProtoMessage() {}

func (x *GetMarkdownContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarkdownContentResponse.ProtoReflect.Descriptor instead.
func (*GetMarkdownContentResponse) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{1}
}

func (x *GetMarkdownContentResponse) GetMarkdown() *models.Markdown {
	if x != nil {
		return x.Markdown
	}
	return nil
}

func (x *GetMarkdownContentResponse) GetPrev() string {
	if x != nil {
		return x.Prev
	}
	return ""
}

func (x *GetMarkdownContentResponse) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

func (x *GetMarkdownContentResponse) GetFooters() []*models.FooterSocial {
	if x != nil {
		return x.Footers
	}
	return nil
}

func (x *GetMarkdownContentResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *GetMarkdownContentResponse) GetThemeColor() string {
	if x != nil {
		return x.ThemeColor
	}
	return ""
}

// 2.GetMarkdownImage
type GetMarkdownImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	RepoName string `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	FilePath string `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *GetMarkdownImageRequest) Reset() {
	*x = GetMarkdownImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarkdownImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarkdownImageRequest) ProtoMessage() {}

func (x *GetMarkdownImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarkdownImageRequest.ProtoReflect.Descriptor instead.
func (*GetMarkdownImageRequest) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{2}
}

func (x *GetMarkdownImageRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetMarkdownImageRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *GetMarkdownImageRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type GetMarkdownImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *GetMarkdownImageResponse) Reset() {
	*x = GetMarkdownImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMarkdownImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMarkdownImageResponse) ProtoMessage() {}

func (x *GetMarkdownImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMarkdownImageResponse.ProtoReflect.Descriptor instead.
func (*GetMarkdownImageResponse) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{3}
}

func (x *GetMarkdownImageResponse) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

// 3.QueryRepoMarkdown
type QueryRepoMarkdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	RepoName       string `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	PlainToTsquery string `protobuf:"bytes,3,opt,name=plain_to_tsquery,json=plainToTsquery,proto3" json:"plain_to_tsquery,omitempty"`
	PageId         int32  `protobuf:"varint,4,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	PageSize       int32  `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *QueryRepoMarkdownRequest) Reset() {
	*x = QueryRepoMarkdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRepoMarkdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRepoMarkdownRequest) ProtoMessage() {}

func (x *QueryRepoMarkdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRepoMarkdownRequest.ProtoReflect.Descriptor instead.
func (*QueryRepoMarkdownRequest) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{4}
}

func (x *QueryRepoMarkdownRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *QueryRepoMarkdownRequest) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *QueryRepoMarkdownRequest) GetPlainToTsquery() string {
	if x != nil {
		return x.PlainToTsquery
	}
	return ""
}

func (x *QueryRepoMarkdownRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *QueryRepoMarkdownRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryRepoMarkdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*models.Markdown `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *QueryRepoMarkdownResponse) Reset() {
	*x = QueryRepoMarkdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRepoMarkdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRepoMarkdownResponse) ProtoMessage() {}

func (x *QueryRepoMarkdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRepoMarkdownResponse.ProtoReflect.Descriptor instead.
func (*QueryRepoMarkdownResponse) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{5}
}

func (x *QueryRepoMarkdownResponse) GetElements() []*models.Markdown {
	if x != nil {
		return x.Elements
	}
	return nil
}

// 4.QueryUserMarkdown
type QueryUserMarkdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PlainToTsquery string `protobuf:"bytes,2,opt,name=plain_to_tsquery,json=plainToTsquery,proto3" json:"plain_to_tsquery,omitempty"`
	PageId         int32  `protobuf:"varint,3,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	PageSize       int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *QueryUserMarkdownRequest) Reset() {
	*x = QueryUserMarkdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserMarkdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserMarkdownRequest) ProtoMessage() {}

func (x *QueryUserMarkdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserMarkdownRequest.ProtoReflect.Descriptor instead.
func (*QueryUserMarkdownRequest) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{6}
}

func (x *QueryUserMarkdownRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *QueryUserMarkdownRequest) GetPlainToTsquery() string {
	if x != nil {
		return x.PlainToTsquery
	}
	return ""
}

func (x *QueryUserMarkdownRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *QueryUserMarkdownRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryUserMarkdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*models.Markdown `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *QueryUserMarkdownResponse) Reset() {
	*x = QueryUserMarkdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserMarkdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserMarkdownResponse) ProtoMessage() {}

func (x *QueryUserMarkdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserMarkdownResponse.ProtoReflect.Descriptor instead.
func (*QueryUserMarkdownResponse) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{7}
}

func (x *QueryUserMarkdownResponse) GetElements() []*models.Markdown {
	if x != nil {
		return x.Elements
	}
	return nil
}

// 4.QueryMarkdown
type QueryMarkdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlainToTsquery string `protobuf:"bytes,1,opt,name=plain_to_tsquery,json=plainToTsquery,proto3" json:"plain_to_tsquery,omitempty"`
	PageId         int32  `protobuf:"varint,2,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	PageSize       int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *QueryMarkdownRequest) Reset() {
	*x = QueryMarkdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMarkdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMarkdownRequest) ProtoMessage() {}

func (x *QueryMarkdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMarkdownRequest.ProtoReflect.Descriptor instead.
func (*QueryMarkdownRequest) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{8}
}

func (x *QueryMarkdownRequest) GetPlainToTsquery() string {
	if x != nil {
		return x.PlainToTsquery
	}
	return ""
}

func (x *QueryMarkdownRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *QueryMarkdownRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryMarkdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*models.Markdown `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *QueryMarkdownResponse) Reset() {
	*x = QueryMarkdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_markdown_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMarkdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMarkdownResponse) ProtoMessage() {}

func (x *QueryMarkdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_markdown_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMarkdownResponse.ProtoReflect.Descriptor instead.
func (*QueryMarkdownResponse) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_markdown_proto_rawDescGZIP(), []int{9}
}

func (x *QueryMarkdownResponse) GetElements() []*models.Markdown {
	if x != nil {
		return x.Elements
	}
	return nil
}

var File_rpcs_rpc_markdown_proto protoreflect.FileDescriptor

var file_rpcs_rpc_markdown_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x64,
	0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x15, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x64, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0xf6, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52,
	0x08, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x65,
	0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x12, 0x2a, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x53, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x52, 0x07, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68, 0x65, 0x6d,
	0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x6f, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x18, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x73, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x54, 0x6f, 0x54, 0x73, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x45, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x4d, 0x61, 0x72,
	0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x08, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x73, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x54, 0x6f, 0x54, 0x73, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x45, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x72,
	0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x08, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x76, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x74, 0x73, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x54, 0x6f, 0x54, 0x73, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x41, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x2f, 0x7a, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x70, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcs_rpc_markdown_proto_rawDescOnce sync.Once
	file_rpcs_rpc_markdown_proto_rawDescData = file_rpcs_rpc_markdown_proto_rawDesc
)

func file_rpcs_rpc_markdown_proto_rawDescGZIP() []byte {
	file_rpcs_rpc_markdown_proto_rawDescOnce.Do(func() {
		file_rpcs_rpc_markdown_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcs_rpc_markdown_proto_rawDescData)
	})
	return file_rpcs_rpc_markdown_proto_rawDescData
}

var file_rpcs_rpc_markdown_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rpcs_rpc_markdown_proto_goTypes = []interface{}{
	(*GetMarkdownContentRequest)(nil),  // 0: pb.GetMarkdownContentRequest
	(*GetMarkdownContentResponse)(nil), // 1: pb.GetMarkdownContentResponse
	(*GetMarkdownImageRequest)(nil),    // 2: pb.GetMarkdownImageRequest
	(*GetMarkdownImageResponse)(nil),   // 3: pb.GetMarkdownImageResponse
	(*QueryRepoMarkdownRequest)(nil),   // 4: pb.QueryRepoMarkdownRequest
	(*QueryRepoMarkdownResponse)(nil),  // 5: pb.QueryRepoMarkdownResponse
	(*QueryUserMarkdownRequest)(nil),   // 6: pb.QueryUserMarkdownRequest
	(*QueryUserMarkdownResponse)(nil),  // 7: pb.QueryUserMarkdownResponse
	(*QueryMarkdownRequest)(nil),       // 8: pb.QueryMarkdownRequest
	(*QueryMarkdownResponse)(nil),      // 9: pb.QueryMarkdownResponse
	(*models.Markdown)(nil),            // 10: pb.Markdown
	(*models.FooterSocial)(nil),        // 11: pb.FooterSocial
	(*timestamppb.Timestamp)(nil),      // 12: google.protobuf.Timestamp
}
var file_rpcs_rpc_markdown_proto_depIdxs = []int32{
	10, // 0: pb.GetMarkdownContentResponse.markdown:type_name -> pb.Markdown
	11, // 1: pb.GetMarkdownContentResponse.footers:type_name -> pb.FooterSocial
	12, // 2: pb.GetMarkdownContentResponse.updated_at:type_name -> google.protobuf.Timestamp
	10, // 3: pb.QueryRepoMarkdownResponse.elements:type_name -> pb.Markdown
	10, // 4: pb.QueryUserMarkdownResponse.elements:type_name -> pb.Markdown
	10, // 5: pb.QueryMarkdownResponse.elements:type_name -> pb.Markdown
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_rpcs_rpc_markdown_proto_init() }
func file_rpcs_rpc_markdown_proto_init() {
	if File_rpcs_rpc_markdown_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcs_rpc_markdown_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarkdownContentRequest); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarkdownContentResponse); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarkdownImageRequest); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMarkdownImageResponse); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRepoMarkdownRequest); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRepoMarkdownResponse); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserMarkdownRequest); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserMarkdownResponse); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMarkdownRequest); i {
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
		file_rpcs_rpc_markdown_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMarkdownResponse); i {
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
			RawDescriptor: file_rpcs_rpc_markdown_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpcs_rpc_markdown_proto_goTypes,
		DependencyIndexes: file_rpcs_rpc_markdown_proto_depIdxs,
		MessageInfos:      file_rpcs_rpc_markdown_proto_msgTypes,
	}.Build()
	File_rpcs_rpc_markdown_proto = out.File
	file_rpcs_rpc_markdown_proto_rawDesc = nil
	file_rpcs_rpc_markdown_proto_goTypes = nil
	file_rpcs_rpc_markdown_proto_depIdxs = nil
}
