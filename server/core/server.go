package core

import (
	"fmt"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 从db加载jwt数据
	if global.GV_DB != nil {
		//system.LoadAll()
	}
	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GV_CONFIG.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GV_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用hust_mall

	项目后台运行地址:http://127.0.0.1%s
	`, address)
	global.GV_LOG.Error(s.ListenAndServe().Error())

}
