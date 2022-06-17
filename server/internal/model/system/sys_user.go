package system

type User struct {
	ID          int64  `json:"ID" form:"ID"`
	UserName    string `json:"username" gorm:"column:username; comment:用户登录名"` // 用户登录名
	Password    string `json:"-"  gorm:"column:password; comment:用户登录密码"`
	Email       string `json:"email" gorm:"column:email; column:email"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`
}

type UserStatisInfo struct {
	ID         int64 `json:"ID" form:"ID"`
	UID        int64 `gorm:"column:uid"`
	OrderCount int   `json:"orderCount" form:"orderCount" gorm:"column:finish_order_count;comment:;"`
	LoginCount int   `json:"loginCount" form:"loginCount" gorm:"column:login_count;comment:;"`
}

// UserLoginLog 结构
type UserLoginLog struct {
	ID        int64  `json:"ID" gorm:"column:id"`
	UID       int64  `json:"uid" gorm:"column:uid"`
	UserName  string `json:"username" form:"username" gorm:"column:username;comment:;"`
	Ip        string `json:"ip" form:"ip" gorm:"column:ip;comment:;"`
	LoginType int    `json:"loginType" form:"loginType" gorm:"column:login_type;comment:;size:1;"`
	City      string `json:"city" form:"city" gorm:"column:city;comment:;"`
	LoginTime int64  `json:"loginTime" form:"loginTime" gorm:"column:login_time;comment:;"`
}

type Online struct {
	Status        int
	LastTimestamp int64
}
