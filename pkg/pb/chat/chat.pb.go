// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: pkg/proto/user/chat.proto

package chat

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

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId    string `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	GroupId     string `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	LastMessage string `protobuf:"bytes,4,opt,name=last_message,json=lastMessage,proto3" json:"last_message,omitempty"`
	Timestamp   int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChatId      int32  `protobuf:"varint,6,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Chat) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

func (x *Chat) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Chat) GetLastMessage() string {
	if x != nil {
		return x.LastMessage
	}
	return ""
}

func (x *Chat) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Chat) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type AddChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId int32 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	GroupId  int32 `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *AddChatRequest) Reset() {
	*x = AddChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChatRequest) ProtoMessage() {}

func (x *AddChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChatRequest.ProtoReflect.Descriptor instead.
func (*AddChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{1}
}

func (x *AddChatRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddChatRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *AddChatRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type AddChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddChatResponse) Reset() {
	*x = AddChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChatResponse) ProtoMessage() {}

func (x *AddChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChatResponse.ProtoReflect.Descriptor instead.
func (*AddChatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{2}
}

func (x *AddChatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  int32  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateChatRequest) Reset() {
	*x = UpdateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChatRequest) ProtoMessage() {}

func (x *UpdateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChatRequest.ProtoReflect.Descriptor instead.
func (*UpdateChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateChatRequest) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *UpdateChatRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int32 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *DeleteChatRequest) Reset() {
	*x = DeleteChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRequest) ProtoMessage() {}

func (x *DeleteChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteChatRequest) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type DeleteChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteChatResponse) Reset() {
	*x = DeleteChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatResponse) ProtoMessage() {}

func (x *DeleteChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatResponse.ProtoReflect.Descriptor instead.
func (*DeleteChatResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteChatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetChatListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetChatListRequest) Reset() {
	*x = GetChatListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatListRequest) ProtoMessage() {}

func (x *GetChatListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatListRequest.ProtoReflect.Descriptor instead.
func (*GetChatListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{6}
}

func (x *GetChatListRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetChatListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*Chat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *GetChatListResponse) Reset() {
	*x = GetChatListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_user_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatListResponse) ProtoMessage() {}

func (x *GetChatListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_user_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatListResponse.ProtoReflect.Descriptor instead.
func (*GetChatListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_user_chat_proto_rawDescGZIP(), []int{7}
}

func (x *GetChatListResponse) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

var File_pkg_proto_user_chat_proto protoreflect.FileDescriptor

var file_pkg_proto_user_chat_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61,
	0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb1, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x32, 0xf8, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x12,
	0x5c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6b, 0x6f, 0x6e,
	0x67, 0x2d, 0x79, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_user_chat_proto_rawDescOnce sync.Once
	file_pkg_proto_user_chat_proto_rawDescData = file_pkg_proto_user_chat_proto_rawDesc
)

func file_pkg_proto_user_chat_proto_rawDescGZIP() []byte {
	file_pkg_proto_user_chat_proto_rawDescOnce.Do(func() {
		file_pkg_proto_user_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_user_chat_proto_rawDescData)
	})
	return file_pkg_proto_user_chat_proto_rawDescData
}

var file_pkg_proto_user_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_user_chat_proto_goTypes = []interface{}{
	(*Chat)(nil),                // 0: chat.Chat
	(*AddChatRequest)(nil),      // 1: chat.AddChatRequest
	(*AddChatResponse)(nil),     // 2: chat.AddChatResponse
	(*UpdateChatRequest)(nil),   // 3: chat.UpdateChatRequest
	(*DeleteChatRequest)(nil),   // 4: chat.DeleteChatRequest
	(*DeleteChatResponse)(nil),  // 5: chat.DeleteChatResponse
	(*GetChatListRequest)(nil),  // 6: chat.GetChatListRequest
	(*GetChatListResponse)(nil), // 7: chat.GetChatListResponse
}
var file_pkg_proto_user_chat_proto_depIdxs = []int32{
	0, // 0: chat.GetChatListResponse.chats:type_name -> chat.Chat
	1, // 1: chat.ChatService.AddChat:input_type -> chat.AddChatRequest
	3, // 2: chat.ChatService.UpdateChat:input_type -> chat.UpdateChatRequest
	4, // 3: chat.ChatService.DeleteChat:input_type -> chat.DeleteChatRequest
	6, // 4: chat.ChatService.GetChatList:input_type -> chat.GetChatListRequest
	2, // 5: chat.ChatService.AddChat:output_type -> chat.AddChatResponse
	2, // 6: chat.ChatService.UpdateChat:output_type -> chat.AddChatResponse
	5, // 7: chat.ChatService.DeleteChat:output_type -> chat.DeleteChatResponse
	7, // 8: chat.ChatService.GetChatList:output_type -> chat.GetChatListResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_user_chat_proto_init() }
func file_pkg_proto_user_chat_proto_init() {
	if File_pkg_proto_user_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_user_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_pkg_proto_user_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChatRequest); i {
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
		file_pkg_proto_user_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddChatResponse); i {
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
		file_pkg_proto_user_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChatRequest); i {
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
		file_pkg_proto_user_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRequest); i {
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
		file_pkg_proto_user_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatResponse); i {
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
		file_pkg_proto_user_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatListRequest); i {
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
		file_pkg_proto_user_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatListResponse); i {
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
			RawDescriptor: file_pkg_proto_user_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_user_chat_proto_goTypes,
		DependencyIndexes: file_pkg_proto_user_chat_proto_depIdxs,
		MessageInfos:      file_pkg_proto_user_chat_proto_msgTypes,
	}.Build()
	File_pkg_proto_user_chat_proto = out.File
	file_pkg_proto_user_chat_proto_rawDesc = nil
	file_pkg_proto_user_chat_proto_goTypes = nil
	file_pkg_proto_user_chat_proto_depIdxs = nil
}
