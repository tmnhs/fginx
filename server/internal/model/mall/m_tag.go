package mall

type Tag struct {
	ID   int64  `json:"ID" form:"ID"`
	Name string `json:"name" form:"name" gorm:"column:name;comment:;"`
}

type TagAndItem struct {
	ID          int64 `json:"ID" form:"ID"`
	ItemId      int64 `json:"itemId" form:"itemId" gorm:"column:item_id"`
	TagId       int64 `json:"tagId" form:"tagId" gorm:"column:tag_id"`
	Type        int   `json:"type" form:"type" gorm:"column:type"`
	CreatedTime int64 `json:"-" form:"-" gorm:"column:created_time;comment:;"`
}
