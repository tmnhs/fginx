package request

type Order struct {
	Uid     int64  `json:"uid" form:"uid" `
	OrderSn string `json:"orderSn" form:"orderSn"`
	Status  int    `json:"status" form:"status" gorm:"column:status;comment:订单状态：1->已完成；2->进行中;3->已关闭(取消)"`
}
