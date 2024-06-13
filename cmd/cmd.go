package cmd

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web-base/conf"
	"github.com/ZhaoJun-hz/go-web-base/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("Clean")
}
