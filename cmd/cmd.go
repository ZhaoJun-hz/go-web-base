package cmd

import (
	"github.com/ZhaoJun-hz/go-web-base/conf"
	"github.com/ZhaoJun-hz/go-web-base/global"
	"github.com/ZhaoJun-hz/go-web-base/router"
)

func Start() {
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	router.InitRouter()
}

func Clean() {
	global.Logger.Info("Clean")
}
