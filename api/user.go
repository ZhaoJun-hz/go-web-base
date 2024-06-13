package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

func (userApi UserApi) Login(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "Login Success",
	})
}
