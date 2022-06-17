package mall

// Category 结构体
type Category struct {
	ID          int64  `json:"ID" form:"ID"`
	Name        string `json:"name" form:"name" gorm:"column:name;comment:;"`
	ParentId    int64  `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;"`
	Level       *int   `json:"level" form:"level" gorm:"column:level;comment:;"`
	ItemCount   *int   `json:"itemCount" form:"itemCount" gorm:"column:item_count;comment:;"`
	ShowStatus  *bool  `json:"showStatus" form:"showStatus" gorm:"column:show_status;comment:;"`
	Sort        *int   `json:"sort" form:"sort" gorm:"column:sort;comment:;"`
	Icon        string `json:"icon" form:"icon" gorm:"column:icon;comment:;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:;"`
}
