package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-main/internal/api"
	"maxblog-me-main/internal/middleware/interceptor"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine      *gin.Engine
	Handler     *api.Handler
	Interceptor *interceptor.Interceptor
}
