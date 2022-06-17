package request

type Profile struct {
	UID      int64  `json:"uid" form:"uid"`
	UserName string `json:"username" form:"username" `
	Email    string `json:"email" `
	Gender   *int   `json:"gender" form:"gender" `
	Region   int    `json:"region" gorm:"column:region;comment:地址区域，1表示韵苑，2表示沁苑，3表示紫菘，4表示其他"`
	Other    string `json:"other" `
	Unit     string `json:"unit" `
	Phone    string `json:"phone" form:"phone"  `
	QQ       string `json:"qq" form:"qq" `
	Wechat   string `json:"wechat" form:"wechat" `
	Avatar   string `json:"avatar" form:"avatar" `
	BgImage  string `json:"bg_image" form:"bg_image" `
}
