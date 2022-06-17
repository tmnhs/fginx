package v1

import (
	"github.com/tmnhs/fginx/server/internal/api/v1/file"
	"github.com/tmnhs/fginx/server/internal/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	FileApiGroup   file.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
