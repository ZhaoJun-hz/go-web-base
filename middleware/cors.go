package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})
}
