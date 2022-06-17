package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	systemReq "github.com/tmnhs/fginx/server/internal/model/system/request"
	redisClient "github.com/tmnhs/fginx/server/internal/service/redis"
	"strconv"
	"strings"
)

var redisService = redisClient.RedisService{}

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	tokenID := c.Request.Header.Get("x-token")
	tokenSplit := strings.Split(tokenID, ":")
	if len(tokenSplit) < 3 {
		return nil, errors.New("split token error")
	}
	key, _ := strconv.Atoi(tokenSplit[2])
	var token string
	err := redisService.GetFromRedis(context.Background(), fmt.Sprintf(model.KeyToken, key), &token)
	if err != nil {
		global.GV_LOG.Error("从Gin的Context中获取从jwt解析信息失败---redis获取错误")
		return nil, err
	}
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GV_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.ID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户Email
func GetUserEmail(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Email
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Email
	}
}

// 从Gin的Context中获取从jwt解析出来的用户Email
func GetUserUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.UserName
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UserName
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}
