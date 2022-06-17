package request

type TagAndItem struct {
	ItemId int64   `json:"itemId" form:"itemId" gorm:"column:item_id"`
	TagIds []int64 `json:"tagIds" form:"tagIds" gorm:"column:tag_id"`
	Type   int     `json:"type" form:"type" gorm:"column:type"`
}
