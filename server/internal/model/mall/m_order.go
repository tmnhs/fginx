package mall

// Order 结构体
// 如果含有time.Time 请自行import time包
type Order struct {
	ID          int64   `json:"ID" form:"ID"`
	OrderSn     string  `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:;"`
	ReceiverId  int64   `json:"receiverId" form:"receiverId" gorm:"column:receiver_id;comment:;"`
	Type        int     `json:"type" form:"type" gorm:"column:type;comment:;"`
	ItemId      int64   `json:"itemId" form:"itemId" gorm:"column:item_id;comment:;"`
	TotalAmount float64 `json:"totalAmount" form:"totalAmount" gorm:"column:total_amount;comment:;"`
	Address     string  `json:"address" form:"address" gorm:"column:address;comment:;"`
	TradeTime   int64   `json:"tradeTime" form:"tradeTime" gorm:"column:trade_time;comment:;"`
	Status      int     `json:"status" form:"status" gorm:"column:status;comment:订单状态：1->已完成；2->进行中;3->已关闭(取消)"`
	Note        string  `json:"note" form:"note" gorm:"column:note;comment:;"`
	//ConfirmStatus *bool    `json:"confirmStatus" form:"confirmStatus" gorm:"column:confirm_status;comment:;"`
	CreatedTime int64 `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:;"`
	UpdatedTime int64 `json:"updatedTime" form:"updatedTime" gorm:"column:updated_time;comment:;"`
}
