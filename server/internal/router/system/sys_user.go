package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/tmnhs/fginx/server/internal/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("change_password", baseApi.ChangePassword) // 用户修改密码
		userRouter.POST("get_profile", baseApi.GetProfile)         //获取用户详细信息
		userRouter.POST("update_profile", baseApi.UpdateProfile)   //更新用户详细信息
		userRouter.POST("login_out", baseApi.LoginOut)             //退出登录
	}
}
