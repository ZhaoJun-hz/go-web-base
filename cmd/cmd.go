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
	router.InitRouter()
}

func Clean() {
	global.Logger.Info("Clean")
}
