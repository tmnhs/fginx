package global

import (
	"github.com/spf13/viper"
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/tmnhs/fginx/server/pkg/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/tmnhs/fginx/server/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

//定义结构
type DbStruct struct {
	DB       *gorm.DB
	TablePre string
}

//定义全局变量
var (
	GV_DB     *gorm.DB
	GV_DBList map[string]DbStruct
	//搜索引擎
	GV_ES     *elastic.Client
	GV_REDIS  *redis.Client
	GV_CONFIG config.Server
	GV_VP     *viper.Viper
	//GV_LOG    *oplogging.Logger
	GV_LOG                 *zap.Logger
	GV_Timer               timer.Timer = timer.NewTimerTask()
	GV_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.Lock()
	defer lock.Unlock()
	return GV_DBList[dbname].DB
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GV_DBList[dbname]
	if !ok {
		panic("db no init")
	}
	return db.DB
}
