//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"maxblog-me-main/internal/api"
	"maxblog-me-main/internal/conf"
	"maxblog-me-main/internal/core"
	"maxblog-me-main/internal/middleware/interceptor"
	dataHandler "maxblog-me-main/src/handler"
	dataService "maxblog-me-main/src/service"
)

func InitInjector() (*Injector, error) {
	wire.Build(
		conf.InitGinEngine,
		api.APISet,
		interceptor.InterceptorSet,
		core.ResponseSet,
		core.LoggerSet,
		dataHandler.HandlerSet,
		dataService.ServiceSet,
		InjectorSet,
	)
	return new(Injector), nil
}
