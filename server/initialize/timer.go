package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func Timer() {
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
