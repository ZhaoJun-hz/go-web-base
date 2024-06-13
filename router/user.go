package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRouter() {
	RegisterRoutes(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Success",
			})
		})

		userRgAuth := rgAuth.Group("/user")

		userRgAuth.GET("", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "zs"},
					{"id": 2, "name": "lisi"},
				},
			})
		})
		userRgAuth.GET("/:id", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id": 1, "name": "zs",
			})
		})
	})
}
