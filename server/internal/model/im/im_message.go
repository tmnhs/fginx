package model

type MessageModel struct {
	Id    int64 `json:"id"` // 自增长id
	MsgId int64 `json:"msg_id" gorm:"column:msg_id"`

	// 客户端消息ID-UUID
	//ClientMsgId string `json:"client_msg_id" gorm:"column:client_msg_id"`
	FromId     int64  `json:"from_id" gorm:"column:from_id" `
	ToId       int64  `json:"to_id" gorm:"column:to_id"`
	MsgType    int    `json:"msg_type" gorm:"column:msg_type"`
	MsgContent string `json:"msg_content" gorm:"column:msg_content"`
	GroupId    int64  `json:"group_id"`
	// cim.CIMResCode
	// 消息错误码 0：一切正常
	//MsgResCode int `json:"msg_res_code" gorm:"column:msg_res_code"`
	// cim.CIMMessageFeature
	// 消息属性 0：默认 1：离线消息 2：漫游消息 3：同步消息 4：透传消息
	MsgFeature int `json:"msg_feature" gorm:"column:msg_feature"`

	// cim.CIMMsgStatus
	// 消息状态 0：默认 1：收到消息，未读 2：收到消息，已读 3：已删 4：发送中 5：已发送
	// 7：草稿 8：发送取消 9：被对方拒绝，如在黑名单中
	MsgStatus   int   `json:"msg_status" gorm:"column:msg_status"`
	CreatedTime int64 `json:"created_time" gorm:"column:created_time"`
	//UpdatedTime   int64 `json:"updated_time" gorm:"column:updated_time"`
}
