package v1


import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/file"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	FileApiGroup file.FileApiGroup
}

var ApiGroupApp = new(ApiGroup)