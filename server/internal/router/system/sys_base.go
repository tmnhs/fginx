package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/tmnhs/fginx/server/internal/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("register", baseApi.Register)           // 用户注册账号
		baseRouter.POST("login", baseApi.Login)                 //通过邮箱/用户名加密码登录
		baseRouter.POST("captcha", baseApi.Captcha)             //生成验证码
		baseRouter.POST("login_by_email", baseApi.LoginByEmail) //通过邮箱验证登录
		baseRouter.POST("send_email", baseApi.SendEmail)        //发送邮件
	}
	return baseRouter
}
