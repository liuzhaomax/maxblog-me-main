package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-main/internal/middleware/interceptor"
	"maxblog-me-main/src/handler"
)

func RegisterRouter(handler *handler.HDemo, app *gin.Engine, itcpt *interceptor.Interceptor) {
	routerDemo := app.Group("/demo")
	{
		routerDemo.GET("", handler.GetDemos)
		routerDemo.GET("/:id", handler.GetDemoById)
	}
}
