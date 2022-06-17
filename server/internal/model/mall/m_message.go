package mall

type Message struct {
	ID          int64  `json:"ID" form:"ID"`
	MsgType     int    `json:"msgType" form:"msgType" gorm:"column:message_type"`
	FromUid     int64  `json:"fromUid" form:"fromUid" gorm:"column:from_uid"`
	ItemId      int64  `json:"itemId" form:"itemId" gorm:"column:item_id"`
	ToUid       int64  `json:"toUid" form:"toUid" gorm:"column:to_uid"`
	Url         string `json:"url" form:"url" gorm:"column:url"`
	Content     string `json:"content" form:"content" gorm:"column:content"`
	ContentType int    `json:"contentType" from:"contentType" gorm:"column:content_type"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	DeletedTime int64  `json:"deletedTime" form:"deletedTime" gorm:"column:deleted_time;comment:;"`
}

type KafkaMessage struct {
	Avatar       string `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
	FromUserName string `protobuf:"bytes,2,opt,name=fromUserName,proto3" json:"fromUserName,omitempty"`
	From         int64  `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To           int64  `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	Content      string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	ContentType  int32  `protobuf:"varint,6,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Type         string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	MessageType  int32  `protobuf:"varint,8,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Url          string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	FileSuffix   string `protobuf:"bytes,10,opt,name=fileSuffix,proto3" json:"fileSuffix,omitempty"`
	File         []byte `protobuf:"bytes,11,opt,name=file,proto3" json:"file,omitempty"`
}
