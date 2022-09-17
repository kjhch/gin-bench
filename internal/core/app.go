package core

import (
	"github.com/gin-gonic/gin"
	"github.com/kjhch/gin-bench/internal/config"
)

type App struct {
	server *gin.Engine
}

func NewApp(engine *gin.Engine) *App {
	return &App{server: engine}
}

func (app *App) Run() {
	err := app.server.Run(config.Get().Server.Http.Addr)
	if err != nil {
		panic(err)
	}
}
