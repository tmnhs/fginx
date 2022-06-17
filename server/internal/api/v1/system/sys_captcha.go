package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model/common/response"
	systemRes "github.com/tmnhs/fginx/server/internal/model/system/response"
	"github.com/tmnhs/fginx/server/pkg/errcode"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

func (b *BaseApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GV_CONFIG.Captcha.ImgHeight, global.GV_CONFIG.Captcha.ImgWidth, global.GV_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)

	if id, b64s, err := cp.Generate(); err != nil {
		global.GV_LOG.Error("验证码获取失败！", zap.Error(err))
		response.FailWithMessage(errcode.ErrorCaptchaGenerate, "验证码获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.GV_CONFIG.Captcha.KeyLong,
		}, "验证码获取成功", c)
		//store.Get(id,true)
		global.GV_LOG.Debug("获取验证码:", zap.String("captcha", store.Get(id, false)))
	}

}
