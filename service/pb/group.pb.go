// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0--rc1
// source: group.proto

package pb

import (
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

type GroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Gid    int64  `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Uid    int64  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Mute   int64  `protobuf:"varint,4,opt,name=mute,proto3" json:"mute,omitempty"`
	Type   int32  `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *GroupMember) Reset() {
	*x = GroupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMember) ProtoMessage() {}

func (x *GroupMember) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMember.ProtoReflect.Descriptor instead.
func (*GroupMember) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupMember) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupMember) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *GroupMember) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GroupMember) GetMute() int64 {
	if x != nil {
		return x.Mute
	}
	return 0
}

func (x *GroupMember) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GroupMember) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type GetMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*GroupMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *GetMembersResponse) Reset() {
	*x = GetMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembersResponse) ProtoMessage() {}

func (x *GetMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembersResponse.ProtoReflect.Descriptor instead.
func (*GetMembersResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

func (x *GetMembersResponse) GetMembers() []*GroupMember {
	if x != nil {
		return x.Members
	}
	return nil
}

type PutMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid    int64        `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Member *GroupMember `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *PutMemberRequest) Reset() {
	*x = PutMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutMemberRequest) ProtoMessage() {}

func (x *PutMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutMemberRequest.ProtoReflect.Descriptor instead.
func (*PutMemberRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{2}
}

func (x *PutMemberRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *PutMemberRequest) GetMember() *GroupMember {
	if x != nil {
		return x.Member
	}
	return nil
}

type GidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid int64 `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
}

func (x *GidRequest) Reset() {
	*x = GidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GidRequest) ProtoMessage() {}

func (x *GidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GidRequest.ProtoReflect.Descriptor instead.
func (*GidRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{3}
}

func (x *GidRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

type GetCidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid int64 `protobuf:"varint,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *GetCidResponse) Reset() {
	*x = GetCidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCidResponse) ProtoMessage() {}

func (x *GetCidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCidResponse.ProtoReflect.Descriptor instead.
func (*GetCidResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{4}
}

func (x *GetCidResponse) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid      int64          `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Name     string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar   string         `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Owner    int64          `protobuf:"varint,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Mute     bool           `protobuf:"varint,5,opt,name=mute,proto3" json:"mute,omitempty"`
	Notice   string         `protobuf:"bytes,6,opt,name=notice,proto3" json:"notice,omitempty"`
	CreateAt int64          `protobuf:"varint,7,opt,name=createAt,proto3" json:"createAt,omitempty"`
	Members  []*GroupMember `protobuf:"bytes,8,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{5}
}

func (x *Group) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Group) GetOwner() int64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *Group) GetMute() bool {
	if x != nil {
		return x.Mute
	}
	return false
}

func (x *Group) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

func (x *Group) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Group) GetMembers() []*GroupMember {
	if x != nil {
		return x.Members
	}
	return nil
}

type RemoveMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid int64   `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Uid []int64 `protobuf:"varint,2,rep,packed,name=uid,proto3" json:"uid,omitempty"`
}

func (x *RemoveMemberRequest) Reset() {
	*x = RemoveMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMemberRequest) ProtoMessage() {}

func (x *RemoveMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMemberRequest.ProtoReflect.Descriptor instead.
func (*RemoveMemberRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveMemberRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *RemoveMemberRequest) GetUid() []int64 {
	if x != nil {
		return x.Uid
	}
	return nil
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid     int64    `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Message *Message `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{7}
}

func (x *NotifyRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *NotifyRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *NotifyRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type DispatchMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DispatchMessageRequest) Reset() {
	*x = DispatchMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchMessageRequest) ProtoMessage() {}

func (x *DispatchMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchMessageRequest.ProtoReflect.Descriptor instead.
func (*DispatchMessageRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{8}
}

func (x *DispatchMessageRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DispatchMessageRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type AddGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group       `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Cid   int64        `protobuf:"varint,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Owner *GroupMember `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *AddGroupRequest) Reset() {
	*x = AddGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGroupRequest) ProtoMessage() {}

func (x *AddGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGroupRequest.ProtoReflect.Descriptor instead.
func (*AddGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{9}
}

func (x *AddGroupRequest) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *AddGroupRequest) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *AddGroupRequest) GetOwner() *GroupMember {
	if x != nil {
		return x.Owner
	}
	return nil
}

type HasMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid int64 `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *HasMemberRequest) Reset() {
	*x = HasMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasMemberRequest) ProtoMessage() {}

func (x *HasMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasMemberRequest.ProtoReflect.Descriptor instead.
func (*HasMemberRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{10}
}

func (x *HasMemberRequest) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *HasMemberRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type HasMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Has bool `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
}

func (x *HasMemberResponse) Reset() {
	*x = HasMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasMemberResponse) ProtoMessage() {}

func (x *HasMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasMemberResponse.ProtoReflect.Descriptor instead.
func (*HasMemberResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{11}
}

func (x *HasMemberResponse) GetHas() bool {
	if x != nil {
		return x.Has
	}
	return false
}

var File_group_proto protoreflect.FileDescriptor

var file_group_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x67, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x75, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x50, 0x0a, 0x10, 0x50, 0x75,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x0a,
	0x47, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x69, 0x64,
	0x22, 0xd1, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x75,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x22, 0x39, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x5d, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x67, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54,
	0x0a, 0x16, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x71, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x28, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x10, 0x48, 0x61, 0x73, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x25, 0x0a, 0x11, 0x48, 0x61, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x68, 0x61, 0x73, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x50, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData = file_group_proto_rawDesc
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_proto_rawDescData)
	})
	return file_group_proto_rawDescData
}

var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_group_proto_goTypes = []interface{}{
	(*GroupMember)(nil),            // 0: proto.GroupMember
	(*GetMembersResponse)(nil),     // 1: proto.GetMembersResponse
	(*PutMemberRequest)(nil),       // 2: proto.PutMemberRequest
	(*GidRequest)(nil),             // 3: proto.GidRequest
	(*GetCidResponse)(nil),         // 4: proto.GetCidResponse
	(*Group)(nil),                  // 5: proto.Group
	(*RemoveMemberRequest)(nil),    // 6: proto.RemoveMemberRequest
	(*NotifyRequest)(nil),          // 7: proto.NotifyRequest
	(*DispatchMessageRequest)(nil), // 8: proto.DispatchMessageRequest
	(*AddGroupRequest)(nil),        // 9: proto.AddGroupRequest
	(*HasMemberRequest)(nil),       // 10: proto.HasMemberRequest
	(*HasMemberResponse)(nil),      // 11: proto.HasMemberResponse
	(*Message)(nil),                // 12: proto.Message
}
var file_group_proto_depIdxs = []int32{
	0,  // 0: proto.GetMembersResponse.members:type_name -> proto.GroupMember
	0,  // 1: proto.PutMemberRequest.member:type_name -> proto.GroupMember
	0,  // 2: proto.Group.members:type_name -> proto.GroupMember
	12, // 3: proto.NotifyRequest.message:type_name -> proto.Message
	12, // 4: proto.DispatchMessageRequest.message:type_name -> proto.Message
	5,  // 5: proto.AddGroupRequest.group:type_name -> proto.Group
	0,  // 6: proto.AddGroupRequest.owner:type_name -> proto.GroupMember
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMember); i {
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
		file_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMembersResponse); i {
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
		file_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutMemberRequest); i {
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
		file_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GidRequest); i {
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
		file_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCidResponse); i {
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
		file_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMemberRequest); i {
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
		file_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRequest); i {
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
		file_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchMessageRequest); i {
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
		file_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGroupRequest); i {
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
		file_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasMemberRequest); i {
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
		file_group_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasMemberResponse); i {
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
			RawDescriptor: file_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_rawDesc = nil
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
