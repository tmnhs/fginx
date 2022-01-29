package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GV_VP = core.Viper()      // 初始化Viper
	global.GV_LOG = core.Zap()       // 初始化zap日志库
	global.GV_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GV_DB != nil {
		initialize.RegisterTables(global.GV_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GV_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
