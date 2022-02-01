package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//fmt.Println(global.GV_CONFIG.System.UseMultipoint )
	if global.GV_CONFIG.System.UseMultipoint {
		//初始化redis服务
		initialize.Redis()
	}
	// 从db加载jwt数据
	if global.GV_DB != nil {
		system.LoadAll()
	}
	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GV_CONFIG.System.Addr)
	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GV_LOG.Info("server run success on ", zap.String("address", address))

	a := []int{1, 3, 2}

	fmt.Println("===================12121", a[len(a)-1])

	fmt.Printf(`
	欢迎使用
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	`, address)
	global.GV_LOG.Error(s.ListenAndServe().Error())

}
