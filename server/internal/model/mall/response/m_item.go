package response

import "github.com/tmnhs/fginx/server/internal/model/mall"

type IdleItemGorm struct {
	mall.IdleItem
	mall.ImageString
}
