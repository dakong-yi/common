// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: pkg/proto/user/friend.proto

package friend

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Friend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Status   int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Friend) Reset() {
	*x = Friend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Friend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Friend) ProtoMessage() {}

func (x *Friend) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Friend.ProtoReflect.Descriptor instead.
func (*Friend) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{0}
}

func (x *Friend) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Friend) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Friend) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Friend) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Friend) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type AddFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId int32 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *AddFriendRequest) Reset() {
	*x = AddFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendRequest) ProtoMessage() {}

func (x *AddFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendRequest.ProtoReflect.Descriptor instead.
func (*AddFriendRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{1}
}

func (x *AddFriendRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddFriendRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

type AddFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddFriendResponse) Reset() {
	*x = AddFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendResponse) ProtoMessage() {}

func (x *AddFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendResponse.ProtoReflect.Descriptor instead.
func (*AddFriendResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{2}
}

func (x *AddFriendResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request message for accepting friend request
type AcceptFriendRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId int32 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *AcceptFriendRequestRequest) Reset() {
	*x = AcceptFriendRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptFriendRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptFriendRequestRequest) ProtoMessage() {}

func (x *AcceptFriendRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptFriendRequestRequest.ProtoReflect.Descriptor instead.
func (*AcceptFriendRequestRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{3}
}

func (x *AcceptFriendRequestRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AcceptFriendRequestRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

// Response message for accepting friend request
type AcceptFriendRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AcceptFriendRequestResponse) Reset() {
	*x = AcceptFriendRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptFriendRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptFriendRequestResponse) ProtoMessage() {}

func (x *AcceptFriendRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptFriendRequestResponse.ProtoReflect.Descriptor instead.
func (*AcceptFriendRequestResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{4}
}

func (x *AcceptFriendRequestResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request message for blacklisting a friend
type BlockFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId int32 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *BlockFriendRequest) Reset() {
	*x = BlockFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockFriendRequest) ProtoMessage() {}

func (x *BlockFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockFriendRequest.ProtoReflect.Descriptor instead.
func (*BlockFriendRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{5}
}

func (x *BlockFriendRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BlockFriendRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

// Response message for blacklisting a friend
type BlockFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BlockFriendResponse) Reset() {
	*x = BlockFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockFriendResponse) ProtoMessage() {}

func (x *BlockFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockFriendResponse.ProtoReflect.Descriptor instead.
func (*BlockFriendResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{6}
}

func (x *BlockFriendResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId int32 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *DeleteFriendRequest) Reset() {
	*x = DeleteFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendRequest) ProtoMessage() {}

func (x *DeleteFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendRequest.ProtoReflect.Descriptor instead.
func (*DeleteFriendRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFriendRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteFriendRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

type DeleteFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteFriendResponse) Reset() {
	*x = DeleteFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendResponse) ProtoMessage() {}

func (x *DeleteFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendResponse.ProtoReflect.Descriptor instead.
func (*DeleteFriendResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFriendResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetFriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetFriendListRequest) Reset() {
	*x = GetFriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendListRequest) ProtoMessage() {}

func (x *GetFriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendListRequest.ProtoReflect.Descriptor instead.
func (*GetFriendListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{9}
}

func (x *GetFriendListRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetFriendListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends []*Friend `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *GetFriendListResponse) Reset() {
	*x = GetFriendListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendListResponse) ProtoMessage() {}

func (x *GetFriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendListResponse.ProtoReflect.Descriptor instead.
func (*GetFriendListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{10}
}

func (x *GetFriendListResponse) GetFriends() []*Friend {
	if x != nil {
		return x.Friends
	}
	return nil
}

type FriendRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends []*Friend `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *FriendRequestList) Reset() {
	*x = FriendRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_friend_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestList) ProtoMessage() {}

func (x *FriendRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_friend_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestList.ProtoReflect.Descriptor instead.
func (*FriendRequestList) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_friend_proto_rawDescGZIP(), []int{11}
}

func (x *FriendRequestList) GetFriends() []*Friend {
	if x != nil {
		return x.Friends
	}
	return nil
}

var File_pkg_proto_user_friend_proto protoreflect.FileDescriptor

var file_pkg_proto_user_friend_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x06, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x52, 0x0a, 0x1a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x1b, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x4a, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4b, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x22,
	0x3d, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x32, 0xd7,
	0x05, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x58, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x2e,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x93, 0x01, 0x0a, 0x13, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x12, 0x7a, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12,
	0x1a, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x74, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x2a,
	0x21, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x77, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6b, 0x6f, 0x6e, 0x67, 0x2d, 0x79, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_user_friend_proto_rawDescOnce sync.Once
	file_pkg_proto_user_friend_proto_rawDescData = file_pkg_proto_user_friend_proto_rawDesc
)

func file_pkg_proto_user_friend_proto_rawDescGZIP() []byte {
	file_pkg_proto_user_friend_proto_rawDescOnce.Do(func() {
		file_pkg_proto_user_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_user_friend_proto_rawDescData)
	})
	return file_pkg_proto_user_friend_proto_rawDescData
}

var file_pkg_proto_user_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pkg_proto_user_friend_proto_goTypes = []interface{}{
	(*Friend)(nil),                      // 0: friend.Friend
	(*AddFriendRequest)(nil),            // 1: friend.AddFriendRequest
	(*AddFriendResponse)(nil),           // 2: friend.AddFriendResponse
	(*AcceptFriendRequestRequest)(nil),  // 3: friend.AcceptFriendRequestRequest
	(*AcceptFriendRequestResponse)(nil), // 4: friend.AcceptFriendRequestResponse
	(*BlockFriendRequest)(nil),          // 5: friend.BlockFriendRequest
	(*BlockFriendResponse)(nil),         // 6: friend.BlockFriendResponse
	(*DeleteFriendRequest)(nil),         // 7: friend.DeleteFriendRequest
	(*DeleteFriendResponse)(nil),        // 8: friend.DeleteFriendResponse
	(*GetFriendListRequest)(nil),        // 9: friend.GetFriendListRequest
	(*GetFriendListResponse)(nil),       // 10: friend.GetFriendListResponse
	(*FriendRequestList)(nil),           // 11: friend.FriendRequestList
}
var file_pkg_proto_user_friend_proto_depIdxs = []int32{
	0,  // 0: friend.GetFriendListResponse.friends:type_name -> friend.Friend
	0,  // 1: friend.FriendRequestList.friends:type_name -> friend.Friend
	1,  // 2: friend.FriendService.AddFriend:input_type -> friend.AddFriendRequest
	3,  // 3: friend.FriendService.AcceptFriendRequest:input_type -> friend.AcceptFriendRequestRequest
	5,  // 4: friend.FriendService.BlockFriend:input_type -> friend.BlockFriendRequest
	7,  // 5: friend.FriendService.DeleteFriend:input_type -> friend.DeleteFriendRequest
	9,  // 6: friend.FriendService.GetFriendList:input_type -> friend.GetFriendListRequest
	9,  // 7: friend.FriendService.GetFriendRequestList:input_type -> friend.GetFriendListRequest
	2,  // 8: friend.FriendService.AddFriend:output_type -> friend.AddFriendResponse
	4,  // 9: friend.FriendService.AcceptFriendRequest:output_type -> friend.AcceptFriendRequestResponse
	6,  // 10: friend.FriendService.BlockFriend:output_type -> friend.BlockFriendResponse
	8,  // 11: friend.FriendService.DeleteFriend:output_type -> friend.DeleteFriendResponse
	10, // 12: friend.FriendService.GetFriendList:output_type -> friend.GetFriendListResponse
	11, // 13: friend.FriendService.GetFriendRequestList:output_type -> friend.FriendRequestList
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_user_friend_proto_init() }
func file_pkg_proto_user_friend_proto_init() {
	if File_pkg_proto_user_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_user_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Friend); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendRequest); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendResponse); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptFriendRequestRequest); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptFriendRequestResponse); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockFriendRequest); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockFriendResponse); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendRequest); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendResponse); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendListRequest); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendListResponse); i {
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
		file_pkg_proto_user_friend_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestList); i {
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
			RawDescriptor: file_pkg_proto_user_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_user_friend_proto_goTypes,
		DependencyIndexes: file_pkg_proto_user_friend_proto_depIdxs,
		MessageInfos:      file_pkg_proto_user_friend_proto_msgTypes,
	}.Build()
	File_pkg_proto_user_friend_proto = out.File
	file_pkg_proto_user_friend_proto_rawDesc = nil
	file_pkg_proto_user_friend_proto_goTypes = nil
	file_pkg_proto_user_friend_proto_depIdxs = nil
}
