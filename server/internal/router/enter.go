package router

import (
	"github.com/tmnhs/fginx/server/internal/router/file"
	"github.com/tmnhs/fginx/server/internal/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	File   file.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
