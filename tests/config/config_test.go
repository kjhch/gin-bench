package config

import (
	"fmt"
	"github.com/kjhch/gin-bench/cmd/wire"
	"github.com/kjhch/gin-bench/internal/config"
	"github.com/kjhch/gin-bench/internal/repo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"testing"
)

func TestGet(t *testing.T) {
	c := config.Get()
	t.Log(config.Get())
	if c.App.Name == "" {
		t.Error("empty conf: gin-bench.name")
	}
}

func TestLoggerFactory(t *testing.T) {
	config.NewLoggerFactory()

	//factory.GetLogger(config.DefaultLogger).Info("default logger msg", zap.String("file", viper.ConfigFileUsed()))
	//factory.GetLogger(config.DefaultLogger).Sugar().Infof("default logger msg")
	//factory.GetLogger(config.DefaultLogger).Sugar().Errorw("default logger msg", "conf", config.Get())
	//
	//factory.GetLogger(config.MetricsLogger).Sugar().Errorw("metrics logger msg", "conf", config.Get())
}

func TestLogger(t *testing.T) {
	eplLogger := zap.NewExample()
	eplLogger.Info("Example Logger", zap.String("config-file", viper.ConfigFileUsed()))

	devLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	devLogger.Info("Dev Logger", zap.String("config-file", viper.ConfigFileUsed()))

	prodLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	prodLogger.Info("Prod Logger", zap.String("config-file", viper.ConfigFileUsed()))
}

func TestDB(t *testing.T) {
	var user = new(repo.User)
	err := config.NewDB().Get(user, "select * from user where id=?", 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

func TestWire(t *testing.T) {
	fmt.Println(wire.InitApp().UserRepo.GetUser(6))

}
