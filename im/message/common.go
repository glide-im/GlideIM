package message

import (
	"encoding/json"
	"fmt"
	"go_im/pkg/logger"
	"strings"
)

const (
	ActionMessage            Action = "message"
	ActionGroupMessage              = "message.group"
	ActionChatMessage               = "message.chat"
	ActionChatMessageRecall         = "message.chat.recall"
	ActionGroupMessageRecall        = "message.group.recall"
	// ActionChatMessageRetry 消息重发, 服务器未ack
	ActionChatMessageRetry = "message.chat.retry"
	// ActionChatMessageResend 消息重发, 服务器已ack, 接收方未ack
	ActionChatMessageResend = "message.chat.resend"
	ActionCSMessage         = "message.cs"
	ActionMessageFailed     = "message.failed.send"

	ActionNotifyNeedAuth      = "notify.auth"
	ActionNotifyKickOut       = "notify.kickout"
	ActionNotifyNewContact    = "notify.contact"
	ActionNotifyGroup         = "notify.group"
	ActionNotifyAccountLogin  = "notify.login"
	ActionNotifyAccountLogout = "notify.logout"
	ActionNotifyError         = "notify.error"

	ActionAckRequest  = "ack.request"
	ActionAckGroupMsg = "ack.group.msg"
	ActionAckMessage  = "ack.message"
	ActionAckNotify   = "ack.notify"

	ActionApi       = "api"
	ActionHeartbeat = "heartbeat"
	ActionApiFailed = "api.failed"
)

type Action string

func (a *Action) Contains(action Action) bool {
	return strings.HasPrefix(string(*a), string(action))
}

type Data struct {
	des interface{}
}

func (d *Data) UnmarshalJSON(bytes []byte) error {
	d.des = bytes
	return nil
}

func (d *Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.des)
}

func (d *Data) Bytes() []byte {
	bytes, ok := d.des.([]byte)
	if ok {
		return bytes
	}
	marshalJSON, err := d.MarshalJSON()
	if err != nil {
		logger.E("message data marshal json error %v", err)
		return nil
	}
	return marshalJSON
}

type Message struct {
	Ver    int64
	Seq    int64
	Action Action
	Data   Data
}

func (m *Message) DeserializeData(v interface{}) error {
	s, ok := m.Data.des.([]byte)
	if ok {
		return json.Unmarshal(s, v)
	}
	return nil
}

func (m *Message) String() string {
	return fmt.Sprintf("Message{Seq=%d, Action=%s, Data=%s}", m.Seq, m.Action)
}

func NewMessage(seq int64, action Action, data interface{}) *Message {
	ret := new(Message)
	ret.Seq = seq
	ret.Action = action
	ret.Data.des = data
	return ret
}
