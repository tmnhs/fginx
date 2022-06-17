package system

import "github.com/tmnhs/fginx/server/internal/service"

type ApiGroup struct {
	BaseApi
	JwtApi
}

var (
	jwtService   = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService  = service.ServiceGroupApp.SystemServiceGroup.UserService
	redisService = service.ServiceGroupApp.RedisServiceGroup.RedisService
)
