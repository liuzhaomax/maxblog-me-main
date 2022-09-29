package api

import (
	"github.com/gin-gonic/gin"
	mw "maxblog-me-main/internal/middleware"
	"maxblog-me-main/internal/middleware/interceptor"
	demoRouter "maxblog-me-main/src/router"
	"net/http"
)

func (handler *Handler) RegisterRouter(app *gin.Engine, itcpt *interceptor.Interceptor) {
	app.NoRoute(handler.GetNoRoute)
	app.Use(mw.Cors())
	app.StaticFS("/static", http.Dir("./static"))
	demoRouter.RegisterRouter(handler.HandlerDemo, app, itcpt)
}

func (handler *Handler) GetNoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"res": "404"})
}
