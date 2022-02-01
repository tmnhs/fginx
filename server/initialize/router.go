package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

//todo
func Routers() *gin.Engine {
	Router := gin.Default()

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /j
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	//Router.StaticFS(global.GV_CONFIG.Local.Path,http.Dir(global.GV_CONFIG.Local.Path))
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GV_LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.GV_LOG.Info("use middleware cors")
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GV_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	// 获取路由组实例
	systemRouter := router.RouterGroupApp.System
	fileRouter := router.RouterGroupApp.File
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitJwtRouter(PrivateGroup)  // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup) // 注册用户路由

		fileRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		// 注册用户路由
	}
	return Router
}
