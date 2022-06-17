package system

import (
	"github.com/tmnhs/fginx/server/global"
)

type JwtBlacklist struct {
	global.GV_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
