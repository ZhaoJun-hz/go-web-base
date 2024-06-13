package cmd

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web-base/conf"
	"github.com/ZhaoJun-hz/go-web-base/global"
	"github.com/ZhaoJun-hz/go-web-base/router"
	"go.uber.org/zap"
)

func Start() {
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	db, err := conf.InitDB()
	if err != nil {
		global.Logger.Error("init db error", zap.Error(err))
		panic(fmt.Sprintf("init db error: %s\n", err.Error()))
	}
	global.DB = db
	redisClient, err := conf.InitRedis()
	if err != nil {
		global.Logger.Error("init redis client error", zap.Error(err))
		panic(fmt.Sprintf("init redis client error: %s\n", err.Error()))
	}
	global.RedisClient = redisClient
	router.InitRouter()
}

func Clean() {
	global.Logger.Info("Clean")
}
