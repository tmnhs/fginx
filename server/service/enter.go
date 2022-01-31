package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/file"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	FileServiceGruop file.FileServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
