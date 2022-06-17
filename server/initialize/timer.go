package initialize

import (
	"fmt"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/config"
	"github.com/tmnhs/fginx/server/pkg/utils"
)

func Timer() {
	//清除数据库
	//ClearDBTask()
	//ClearTimeoutConnections()
}
func ClearDBTask() {
	if global.GV_CONFIG.Timer.Start {
		for i := range global.GV_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GV_Timer.AddTaskByFunc("ClearDB", global.GV_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GV_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error: ", err)
					}
				})
			}(global.GV_CONFIG.Timer.Detail[i])
		}
	}
}
