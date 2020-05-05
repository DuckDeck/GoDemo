package initialize

import (
	"GoDemo/Project/Mall/api"
	"GoDemo/Project/Mall/global"
	middleware "GoDemo/Project/Mall/middlemare"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.Cors())
	global.G_LOG.Debug("use middleware cors")
	ApiGroup := Router.Group("")
	initMainRouter(ApiGroup)
	global.G_LOG.Debug("router register success")
	return Router
}

func initMainRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("main")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.GET("/", api.Main) // 修改密码
		UserRouter.GET("/list", api.Main)
	}
}
