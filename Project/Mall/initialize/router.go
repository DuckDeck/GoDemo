package initialize

import (
	"GoDemo/Project/Mall/api"
	"GoDemo/Project/Mall/global"
	middleware "GoDemo/Project/Mall/middlemare"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.Cors())
	Router.StaticFS("/file", http.Dir("./static/img"))
	global.G_LOG.Debug("use middleware cors")
	ApiGroup := Router.Group("")
	initMainRouter(ApiGroup)
	uploadRouter(ApiGroup)
	global.G_LOG.Debug("router register success")
	return Router
}

func initMainRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("main")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.GET("/", api.Main)
	}
}

func uploadRouter(Router *gin.RouterGroup) {
	UploadRouter := Router.Group("upload")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UploadRouter.POST("/header", api.Upload)
		UploadRouter.POST("/img", api.Upload)
	}
}
