package mall

type Address struct {
	ID     int64  `json:"ID" gorm:"column:id"`
	Region int    `json:"region" gorm:"column:region;comment:地址区域，1表示韵苑，2表示沁苑，3表示紫菘，4表示其他"`
	Other  string `json:"other" gorm:"column:other"`
	Unit   string `json:"unit" gorm:"column:unit"`
}
