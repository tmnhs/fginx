package mall

// Feedback 结构体
// 如果含有time.Time 请自行import time包
type Feedback struct {
	ID          int64  `json:"ID" form:"ID"`
	Uid         int64  `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	Content     string `json:"content" form:"content" gorm:"column:content;comment:;"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
}
