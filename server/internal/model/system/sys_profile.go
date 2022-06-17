package system

type ProfileUpdate struct {
	UID         int64  `json:"uid" form:"uid" gorm:"column:uid"`
	UserName    string `json:"username" form:"username" gorm:"column:username;comment:;"`
	Email       string `json:"email" gorm:"column:email"`
	Gender      *int   `json:"gender" form:"gender" gorm:"column:gender;comment:;size:1;"`
	AddressID   int64  `gorm:"column:address_id;comment:;size:1;"`
	Phone       string `json:"phone" form:"phone" gorm:"column:phone" `
	QQ          string `json:"qq" form:"qq" gorm:"column:qq"`
	Wechat      string `json:"wechat" form:"wechat" gorm:"column:wechat"`
	Avatar      string `json:"avatar" form:"avatar" gorm:"column:avatar"`
	BgImage     string `json:"bg_image" form:"bg_image" gorm:"column:bg_image"`
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`
}

type Profile struct {
	UID         int64  `json:"uid" form:"uid" gorm:"column:uid"`
	UserName    string `json:"username" form:"username" gorm:"column:username;comment:;"`
	Email       string `json:"email" gorm:"column:email"`
	Gender      *int   `json:"gender" form:"gender" gorm:"column:gender;comment:;size:1;"`
	AddressID   int64  `gorm:"column:address_id;comment:;size:1;"`
	IsOnline    int    `json:"isOnline" form:"isOnline" gorm:"column:is_online;comment:;size:1;"`
	Status      int    `json:"-" form:"-" gorm:"column:status"`
	Phone       string `json:"phone" form:"phone" gorm:"column:phone" `
	QQ          string `json:"qq" form:"qq" gorm:"column:qq"`
	Wechat      string `json:"wechat" form:"wechat" gorm:"column:wechat"`
	Avatar      string `json:"avatar" form:"avatar" gorm:"column:avatar"`
	BgImage     string `json:"bg_image" form:"bg_image" gorm:"column:bg_image"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`
}
