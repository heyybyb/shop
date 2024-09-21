package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/middlewares"
	"shop-api/user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors()) //配置跨域
	apiGroup := Router.Group("/u/v1")
	router.InitUserRouter(apiGroup)
	router.InitBaseRouter(apiGroup)
	return Router
}
