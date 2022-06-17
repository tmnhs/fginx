package response

import (
	"github.com/tmnhs/fginx/server/internal/model/mall"
	"github.com/tmnhs/fginx/server/internal/model/system"
)

type ProfileResponse struct {
	Address *mall.Address   `json:"address"`
	Profile *system.Profile `json:"profile" `
}
