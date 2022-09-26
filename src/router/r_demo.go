package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-main/src/handler"
)

func RegisterRouter(handler *handler.HDemo, group *gin.RouterGroup) {
	routerDemo := group.Group("/demo")
	{
		routerDemo.GET("", handler.GetDemos)
		routerDemo.GET("/:id", handler.GetDemoById)
	}
}
