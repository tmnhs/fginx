package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type JwtBlacklist struct {
	global.GV_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
