// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: models/user.proto

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

type UserBasicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username   string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email      string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Motto      string                 `protobuf:"bytes,4,opt,name=motto,proto3" json:"motto,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Onboarding bool                   `protobuf:"varint,8,opt,name=onboarding,proto3" json:"onboarding,omitempty"`
}

func (x *UserBasicInfo) Reset() {
	*x = UserBasicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBasicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBasicInfo) ProtoMessage() {}

func (x *UserBasicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBasicInfo.ProtoReflect.Descriptor instead.
func (*UserBasicInfo) Descriptor() ([]byte, []int) {
	return file_models_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserBasicInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserBasicInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserBasicInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserBasicInfo) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *UserBasicInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserBasicInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserBasicInfo) GetOnboarding() bool {
	if x != nil {
		return x.Onboarding
	}
	return false
}

type UserImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Avatar            []byte                 `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	UpdateImageInfoAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_image_info_at,json=updateImageInfoAt,proto3" json:"update_image_info_at,omitempty"`
}

func (x *UserImageInfo) Reset() {
	*x = UserImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserImageInfo) ProtoMessage() {}

func (x *UserImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserImageInfo.ProtoReflect.Descriptor instead.
func (*UserImageInfo) Descriptor() ([]byte, []int) {
	return file_models_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserImageInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserImageInfo) GetAvatar() []byte {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *UserImageInfo) GetUpdateImageInfoAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateImageInfoAt
	}
	return nil
}

type DailyCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Date  string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *DailyCount) Reset() {
	*x = DailyCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyCount) ProtoMessage() {}

func (x *DailyCount) ProtoReflect() protoreflect.Message {
	mi := &file_models_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyCount.ProtoReflect.Descriptor instead.
func (*DailyCount) Descriptor() ([]byte, []int) {
	return file_models_user_proto_rawDescGZIP(), []int{2}
}

func (x *DailyCount) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DailyCount) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type UserCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Count    int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UserCount) Reset() {
	*x = UserCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCount) ProtoMessage() {}

func (x *UserCount) ProtoReflect() protoreflect.Message {
	mi := &file_models_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCount.ProtoReflect.Descriptor instead.
func (*UserCount) Descriptor() ([]byte, []int) {
	return file_models_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserCount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCount) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UserCountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CountLikes     int32 `protobuf:"varint,2,opt,name=count_likes,json=countLikes,proto3" json:"count_likes,omitempty"`
	CountFollowing int32 `protobuf:"varint,3,opt,name=count_following,json=countFollowing,proto3" json:"count_following,omitempty"`
	CountFollower  int32 `protobuf:"varint,4,opt,name=count_follower,json=countFollower,proto3" json:"count_follower,omitempty"`
	CountRepos     int32 `protobuf:"varint,5,opt,name=count_repos,json=countRepos,proto3" json:"count_repos,omitempty"`
	Following      bool  `protobuf:"varint,6,opt,name=following,proto3" json:"following,omitempty"`
}

func (x *UserCountInfo) Reset() {
	*x = UserCountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCountInfo) ProtoMessage() {}

func (x *UserCountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCountInfo.ProtoReflect.Descriptor instead.
func (*UserCountInfo) Descriptor() ([]byte, []int) {
	return file_models_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserCountInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserCountInfo) GetCountLikes() int32 {
	if x != nil {
		return x.CountLikes
	}
	return 0
}

func (x *UserCountInfo) GetCountFollowing() int32 {
	if x != nil {
		return x.CountFollowing
	}
	return 0
}

func (x *UserCountInfo) GetCountFollower() int32 {
	if x != nil {
		return x.CountFollower
	}
	return 0
}

func (x *UserCountInfo) GetCountRepos() int32 {
	if x != nil {
		return x.CountRepos
	}
	return 0
}

func (x *UserCountInfo) GetFollowing() bool {
	if x != nil {
		return x.Following
	}
	return false
}

type ListUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email      string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Blocked    bool                   `protobuf:"varint,3,opt,name=blocked,proto3" json:"blocked,omitempty"`
	Verified   bool                   `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
	Onboarding bool                   `protobuf:"varint,6,opt,name=onboarding,proto3" json:"onboarding,omitempty"`
	Role       string                 `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ListUserInfo) Reset() {
	*x = ListUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserInfo) ProtoMessage() {}

func (x *ListUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_models_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserInfo.ProtoReflect.Descriptor instead.
func (*ListUserInfo) Descriptor() ([]byte, []int) {
	return file_models_user_proto_rawDescGZIP(), []int{5}
}

func (x *ListUserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListUserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ListUserInfo) GetBlocked() bool {
	if x != nil {
		return x.Blocked
	}
	return false
}

func (x *ListUserInfo) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *ListUserInfo) GetOnboarding() bool {
	if x != nil {
		return x.Onboarding
	}
	return false
}

func (x *ListUserInfo) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ListUserInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ListUserInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_models_user_proto protoreflect.FileDescriptor

var file_models_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x4b, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x41,
	0x74, 0x22, 0x36, 0x0a, 0x0a, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3d, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x22, 0xa0, 0x02, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x7a, 0x64, 0x6c, 0x70, 0x2f, 0x7a, 0x62, 0x6f, 0x6f,
	0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_models_user_proto_rawDescOnce sync.Once
	file_models_user_proto_rawDescData = file_models_user_proto_rawDesc
)

func file_models_user_proto_rawDescGZIP() []byte {
	file_models_user_proto_rawDescOnce.Do(func() {
		file_models_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_user_proto_rawDescData)
	})
	return file_models_user_proto_rawDescData
}

var file_models_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_models_user_proto_goTypes = []interface{}{
	(*UserBasicInfo)(nil),         // 0: pb.UserBasicInfo
	(*UserImageInfo)(nil),         // 1: pb.UserImageInfo
	(*DailyCount)(nil),            // 2: pb.DailyCount
	(*UserCount)(nil),             // 3: pb.UserCount
	(*UserCountInfo)(nil),         // 4: pb.UserCountInfo
	(*ListUserInfo)(nil),          // 5: pb.ListUserInfo
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_models_user_proto_depIdxs = []int32{
	6, // 0: pb.UserBasicInfo.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: pb.UserBasicInfo.updated_at:type_name -> google.protobuf.Timestamp
	6, // 2: pb.UserImageInfo.update_image_info_at:type_name -> google.protobuf.Timestamp
	6, // 3: pb.ListUserInfo.updated_at:type_name -> google.protobuf.Timestamp
	6, // 4: pb.ListUserInfo.created_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_models_user_proto_init() }
func file_models_user_proto_init() {
	if File_models_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBasicInfo); i {
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
		file_models_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserImageInfo); i {
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
		file_models_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyCount); i {
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
		file_models_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCount); i {
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
		file_models_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCountInfo); i {
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
		file_models_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserInfo); i {
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
			RawDescriptor: file_models_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_user_proto_goTypes,
		DependencyIndexes: file_models_user_proto_depIdxs,
		MessageInfos:      file_models_user_proto_msgTypes,
	}.Build()
	File_models_user_proto = out.File
	file_models_user_proto_rawDesc = nil
	file_models_user_proto_goTypes = nil
	file_models_user_proto_depIdxs = nil
}
