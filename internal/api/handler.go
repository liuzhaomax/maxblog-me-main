package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-main/internal/middleware/interceptor"
	dataHandler "maxblog-me-main/src/handler"
)

var APISet = wire.NewSet(wire.Struct(new(Handler), "*"), wire.Bind(new(IHandler), new(*Handler)))

type Handler struct {
	HandlerDemo *dataHandler.HDemo
}

type IHandler interface {
	Register(app *gin.Engine, itcpt *interceptor.Interceptor)
}

func (handler *Handler) Register(app *gin.Engine, itcpt *interceptor.Interceptor) {
	handler.RegisterRouter(app, itcpt)
}
