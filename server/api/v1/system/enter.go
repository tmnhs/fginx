package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BaseApi
	JwtApi
}

var (
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
)
