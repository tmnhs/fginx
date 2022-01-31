package router
import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/file"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  file.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
