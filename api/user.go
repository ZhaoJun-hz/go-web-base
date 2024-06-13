package api

import (
	"github.com/ZhaoJun-hz/go-web-base/service"
	"github.com/ZhaoJun-hz/go-web-base/service/dto"
	"github.com/ZhaoJun-hz/go-web-base/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

func (userApi UserApi) Login(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO
	if err := userApi.BuildRequest(BuildRequestOption{
		Ctx: ctx,
		DTO: &userLoginDTO,
	}).GetError(); err != nil {
		return
	}
	user, err := userApi.Service.Login(userLoginDTO)
	if err != nil {
		userApi.Fail(Response{
			Msg: err.Error(),
		})
		return
	}
	token, _ := utils.GeneratorToken(user.ID)
	userApi.OK(Response{
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}
