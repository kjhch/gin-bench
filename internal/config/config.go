package config

import "time"

var appConf = new(AppConfig)

type AppConfig struct {
	App        App
	Server     Server
	DataSource DataSource
	Logger     map[string]LoggerConf
}

type App struct {
	Name   string
	Author string
	Since  time.Time
}

type Server struct {
	Http struct {
		Addr string
	}
}

type DataSource struct {
	Mysql struct {
		User           string
		Password       string
		Addr           string
		DBName         string
		MaxOpenConnNum int
	}
	Redis struct {
		Addr     string
		Password string
		PoolSize int
	}
}

func Get() *AppConfig {
	return appConf
}
