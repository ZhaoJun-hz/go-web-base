package router

import (
	"github.com/ZhaoJun-hz/go-web-base/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter() {
	RegisterRoutes(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		userRgPublic := rgPublic.Group("/user")
		{
			userRgPublic.POST("/login", userApi.Login)
		}

		userRgAuth := rgAuth.Group("/user")
		{
			userRgAuth.POST("", userApi.AddUser)
			userRgAuth.GET("/:id", userApi.GetUserInfo)
			userRgAuth.POST("/list", userApi.GetUserList)
			userRgAuth.PUT("", userApi.UpdateUser)
			userRgAuth.DELETE("/:id", userApi.DeleteUser)
		}
	})
}
