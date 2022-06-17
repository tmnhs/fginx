package core

import (
	"flag"
	"fmt"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/pkg/utils"
	"os"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else if config == utils.Development {
			config = utils.LocalConfig
			fmt.Printf("您正在使用命令行的-c参数传递的值,development为本地开发环境,config的路径为%v\n", config)
		} else if config == utils.Production {
			config = utils.ProductConfig
			fmt.Printf("您正在使用命令行的-c参数传递的值,production为生产环境,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GV_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GV_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.GV_CONFIG.JWT.ExpiresTime)),
	)
	return v
}
