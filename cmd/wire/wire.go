//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kjhch/gin-bench/internal/api"
	"github.com/kjhch/gin-bench/internal/config"
	"github.com/kjhch/gin-bench/internal/core"
	"github.com/kjhch/gin-bench/internal/repo"
	"github.com/kjhch/gin-bench/internal/service"
)

var confSet = wire.NewSet(config.NewLoggerFactory, config.NewDB, config.NewRDB)

var coreSet = wire.NewSet(core.NewHttpServer, core.NewApp)

var apiSet = wire.NewSet(api.NewUserApi)

var serviceSet = wire.NewSet(service.NewUserService)

var repoSet = wire.NewSet(repo.NewUserRepo)

func InitApp() *core.App {
	panic(wire.Build(confSet, coreSet, apiSet, serviceSet, repoSet))
}
