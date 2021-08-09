package entity

import "go_im/im/dao"

type LoginRequest struct {
	Device   int64  `json:"device"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Token    string `json:"token"`
	DeviceId int64  `json:"device_id"`
	Uid      int64  `json:"uid"`
}

type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// AuthorResponse login or register result
type AuthorResponse struct {
	Token string
	Uid   int64
}

type UserInfoRequest struct {
	Uid []int64
}

type UserInfoResponse struct {
	Uid      int64
	Nickname string
	Avatar   string
}

type UserInfoListResponse struct {
	UserInfo []*UserInfoResponse
}

type UserNewChatRequest struct {
	Id   int64
	Type int8
}

type ContactResponse struct {
	Id     int64
	Avatar string
	Name   string
	Type   int8
}

type ChatHistoryRequest struct {
	Cid  int64
	Time int64
	Type int8
}

type ChatInfoRequest struct {
	Cid int64
}

type GroupInfoRequest struct {
	Gid []int64
}

type CreateGroupRequest struct {
	Name   string
	Member []int64
}

type GroupResponse struct {
	dao.Group
	Members []*dao.GroupMember
}

type AddedGroupResponse struct {
	Group     *dao.Group
	GroupChat *dao.Chat
}

type JoinGroupRequest struct {
	Gid int64
}

type ExitGroupRequest struct {
	Gid int64
}

type GetGroupMemberRequest struct {
	Gid int64
}

type GroupMemberResponse struct {
	Uid        int64
	Nickname   string
	RemarkName string
	Type       int8
	Online     bool
	Mute       bool
}

type AddMemberRequest struct {
	Gid int64
	Uid []int64
}

type AddFriendRequest struct {
	Uid    int64
	Remark string
}

type RemoveMemberRequest struct {
	Gid int64
	Uid []int64
}
