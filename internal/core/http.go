package core

import (
	"github.com/gin-gonic/gin"
	"github.com/kjhch/gin-bench/internal/api"
	"github.com/kjhch/gin-bench/internal/config"
	"net/http"
)

func NewHttpServer(userApi *api.UserApi) *gin.Engine {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"gin-bench": config.Get().App.Name})
	})
	engine.Handle(http.MethodGet, "/user", userApi.GetUser)
	engine.Handle(http.MethodGet, "/users", userApi.ListUsers)
	return engine
}
