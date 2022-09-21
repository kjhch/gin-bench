package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewDB() *sqlx.DB {
	mysqlConf := mysql.NewConfig()
	mysqlConf.User = appConf.DataSource.Mysql.User
	mysqlConf.Passwd = appConf.DataSource.Mysql.Password
	mysqlConf.Net = "tcp"
	mysqlConf.Addr = appConf.DataSource.Mysql.Addr
	mysqlConf.DBName = appConf.DataSource.Mysql.DBName
	mysqlConf.ParseTime = true
	mysqlConf.Loc = time.Local

	fmt.Println("connect to ", mysqlConf.FormatDSN())
	db := sqlx.MustConnect("mysql", mysqlConf.FormatDSN())
	db.SetMaxOpenConns(appConf.DataSource.Mysql.MaxOpenConnNum)
	db.SetMaxIdleConns(appConf.DataSource.Mysql.MaxOpenConnNum)
	return db
}

func NewRDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     appConf.DataSource.Redis.Addr,
		Password: appConf.DataSource.Redis.Password,
		PoolSize: appConf.DataSource.Redis.PoolSize,
	})
}
