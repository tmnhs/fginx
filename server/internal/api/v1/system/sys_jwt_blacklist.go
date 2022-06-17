package system

import (
	"github.com/gin-gonic/gin"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model/common/response"
	"github.com/tmnhs/fginx/server/internal/model/system"
	"github.com/tmnhs/fginx/server/pkg/errcode"
	"go.uber.org/zap"
)

type JwtApi struct{}

func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JoinInBlacklist(jwt); err != nil {
		global.GV_LOG.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage(errcode.ErrorJwtInvalid, "jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
