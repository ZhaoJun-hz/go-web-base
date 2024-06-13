package conf

import (
	"fmt"
	"github.com/ZhaoJun-hz/go-web-base/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Load config file error: %s\n", err.Error()))
		panic(fmt.Sprintf("Load config file error: %s\n", err.Error()))
	}

}
