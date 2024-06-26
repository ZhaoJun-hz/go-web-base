package router

import (
	"context"
	"fmt"
	"github.com/ZhaoJun-hz/go-web-base/global"
	"github.com/ZhaoJun-hz/go-web-base/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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
	ctx, cancelContext := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelContext()
	engine := gin.Default()
	engine.Use(middleware.Cors())

	rgPublic := engine.Group("/api/v1/public")
	rgAuth := engine.Group("/api/v1/")
	rgAuth.Use(middleware.Auth())

	InitBasePlatformRoutes()
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "9876"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	go func() {
		global.Logger.Info("start http server on ", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("start http server err: %s", err.Error()))
			return
		}
	}()
	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("http server shutdown err: %s", err.Error()))
		return
	}
	global.Logger.Info(fmt.Sprintf("http server shutdown success"))
}

func InitBasePlatformRoutes() {
	InitUserRouter()
}
