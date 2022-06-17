package main

import (
	"github.com/tmnhs/fginx/server/core"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GV_VP = core.Viper()      // 初始化Viper
	global.GV_LOG = core.Zap()       // 初始化zap日志库
	global.GV_DB = initialize.Gorm() // gorm连接数据库
	initialize.Redis()               //redis连接
	initialize.Timer()
	//initialize.Elastic() //搜索引擎
	//initialize.Kafka()   //kafka消息引擎
	//initialize.DBList()
	if global.GV_DB != nil {
		initialize.RegisterTables(global.GV_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GV_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
