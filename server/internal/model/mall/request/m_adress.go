package request

type Address struct {
	Region int `json:"region" form:"region" gorm:"column:region;comment:地址区域，1表示韵苑，2表示沁苑，3表示紫菘，4表示其他"`
}
