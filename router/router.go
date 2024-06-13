package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRoute
)

func RegisterRoutes(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)

}

func InitRouter() {
	engine := gin.Default()

	rgPublic := engine.Group("/api/v1/public")
	rgAuth := engine.Group("/api/v1/")

	InitBasePlatformRoutes()
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "9876"
	}
	err := engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf("Start server error: %s", err.Error()))
	}
}

func InitBasePlatformRoutes() {
	InitUserRouter()
}
