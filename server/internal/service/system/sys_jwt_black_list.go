package system

import (
	"context"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	"github.com/tmnhs/fginx/server/internal/model/system"
	"go.uber.org/zap"
	"time"
)

type JwtService struct{}

func (jwtService *JwtService) JoinInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GV_DB.Table(model.JwtTableName).Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

func (jetService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GV_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GV_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GV_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GV_DB.Table(model.JwtTableName).Select("jwt").Find(&data).Error
	if err != nil {
		global.GV_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
