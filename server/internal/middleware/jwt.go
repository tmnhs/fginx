package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	"github.com/tmnhs/fginx/server/internal/model/common/response"
	"github.com/tmnhs/fginx/server/internal/model/system"
	"github.com/tmnhs/fginx/server/internal/service"
	"github.com/tmnhs/fginx/server/pkg/errcode"
	"github.com/tmnhs/fginx/server/pkg/timer"
	"github.com/tmnhs/fginx/server/pkg/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var redisService = service.ServiceGroupApp.RedisServiceGroup.RedisService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		tokenID := c.Request.Header.Get("x-token")
		tokenSplit := strings.Split(tokenID, ":")
		if len(tokenSplit) < 3 {
			global.GV_LOG.Debug("get jwt  token error ,you have no right ")
			response.FailWithDetailed(errcode.ErrorNotLogin, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		key, _ := strconv.Atoi(tokenSplit[2])
		var token string
		err := redisService.GetFromRedis(context.Background(), fmt.Sprintf(model.KeyToken, key), &token)
		if err != nil && err != model.ErrRedisNotFound {
			response.FailWithDetailed(errcode.ErrorRedis, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if token == "" {
			global.GV_LOG.Debug("get jwt  token error ,you have no right ")
			response.FailWithDetailed(errcode.ErrorNotLogin, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(errcode.ErrorJwtInvalid, gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(errcode.ErrorTokenExpiration, gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(errcode.ERROR, gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		//if err, _ = userService.FindUserByUuid(claims.UUID.String()); err != nil {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.GV_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.GV_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.UserName)
				if err != nil {
					global.GV_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JoinInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.UserName)
			}
		}
		//设置用户在线状态
		err = userService.SetUserOnline(context.Background(), claims.ID, model.UserOnline)
		if err != nil {
			global.GV_LOG.Error("SetUserOnline err:", zap.Error(err))
			response.FailWithDetailed(errcode.ERROR, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		err = redisService.SetToRedis(context.Background(), fmt.Sprintf(model.KeyLastLogin, claims.ID), timer.GetNowUnix(), 0)
		if err != nil {
			response.FailWithDetailed(errcode.ERROR, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)

		c.Next()
	}
}
