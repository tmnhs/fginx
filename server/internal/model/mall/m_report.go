package mall

// Report 结构体
type Report struct {
	ID          int64  `json:"ID" form:"ID"`
	Uid         int64  `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
	ItemId      int64  `json:"itemId" form:"itemId" gorm:"column:item_id;comment:;"`
	Type        int    `json:"type" form:"type" gorm:"column:type;comment:;"`
	ItemType    int    `json:"itemType" form:"itemType" gorm:"column:item_type;comment:物品类型  1->闲置  2->求购"`
	Other       string `json:"other" form:"other" gorm:"column:other;comment:举报类型type为其他时的举报内容"`
	CreatedTime int64  `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
}
