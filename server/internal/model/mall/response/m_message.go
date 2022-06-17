package response

type Message struct {
	ID           int64  `json:"ID" form:"ID" `
	FromUid      int64  `json:"fromUid" form:"fromUid" gorm:"column:from_uid"`
	ToUid        int64  `json:"toUid" form:"toUid" gorm:"column:to_uid"`
	Content      string `json:"content" form:"content" gorm:"column:content"`
	ContentType  int    `json:"contentType" form:"contentType" gorm:"column:content_type;comment:'消息内容类型：1文字，2语音，3视频'"`
	CreatedTime  int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	FromUsername string `json:"fromUsername" form:"fromUsername" gorm:"-"`
	ToUsername   string `json:"toUsername" form:"toUsername" gorm:"-"`
	Avatar       string `json:"avatar" form:"avatar" gorm:"column:avatar"`
	Url          string `json:"url" form:"url" gorm:"url"`
}
