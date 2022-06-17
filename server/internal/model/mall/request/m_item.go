package request

import (
	"github.com/tmnhs/fginx/server/internal/model/common/request"
	"github.com/tmnhs/fginx/server/internal/model/mall"
)

type ItemSearch struct {
	request.PageInfo
	CategoryId int64  `json:"categoryId" form:"categoryId"`
	Search     string `json:"search" form:"search"`
	PriceSort  int    `json:"priceSort" form:"priceSort"` //是否按价格排序，0->否，1->是
	TimeSort   int    `json:"timeSort" form:"timeSort"`   //是否按时间排序，0->否，1->是
}

type ReqIdleItem struct {
	mall.IdleItem
	mall.ImageArray
}
