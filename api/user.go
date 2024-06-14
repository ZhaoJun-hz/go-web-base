package api

import (
	"context"
	"fmt"
	"github.com/ZhaoJun-hz/go-web-base/global"
	"github.com/ZhaoJun-hz/go-web-base/global/constants"
	"github.com/ZhaoJun-hz/go-web-base/service"
	"github.com/ZhaoJun-hz/go-web-base/service/dto"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

const (
	ERR_CODE_ADD_USER      = 10011
	ERR_CODE_GET_USER_INFO = 10012
	ERR_CODE_GET_USER_LIST = 10013
	ERR_CODE_UPDATE_USER   = 10014
	ERR_CODE_DELETE_USER   = 10015
	ERR_CODE_REGISTER_USER = 10016
	ERR_CODE_LOGIN         = 10017
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

func (m UserApi) Login(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userLoginDTO}).GetError(); err != nil {
		return
	}
	user, token, err := m.Service.Login(userLoginDTO)
	if err == nil {
		global.RedisClient.Set(context.Background(),
			strings.ReplaceAll(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(user.ID))),
			token,
			viper.GetDuration("jwt.tokenExpire")*time.Minute)
	}
	if err != nil {
		m.Fail(Response{
			Code: ERR_CODE_LOGIN,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(Response{
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func (m UserApi) Register(ctx *gin.Context) {
	var dto dto.UserRegisterDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &dto}).GetError(); err != nil {
		return
	}
	err := m.Service.Register(dto)
	if err != nil {
		m.ServerFail(Response{
			Code: ERR_CODE_REGISTER_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(Response{})
}

func (m UserApi) AddUser(ctx *gin.Context) {
	var userAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userAddDTO}).GetError(); err != nil {
		return
	}

	file, _ := ctx.FormFile("file")
	if file != nil {
		filePath := fmt.Sprintf("./upload/%s", file.Filename)
		_ = ctx.SaveUploadedFile(file, filePath)
		userAddDTO.Avatar = filePath
	}
	err := m.Service.AddUser(&userAddDTO)
	if err != nil {
		m.ServerFail(Response{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(Response{
		Data: userAddDTO,
	})
}

func (m UserApi) GetUserInfo(ctx *gin.Context) {
	var dto dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &dto, BindUri: true}).GetError(); err != nil {
		return
	}
	user, err := m.Service.GetUserById(&dto)
	if err != nil {
		m.ServerFail(Response{
			Code: ERR_CODE_GET_USER_INFO,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(Response{
		Data: user,
	})
}

func (m UserApi) GetUserList(ctx *gin.Context) {
	var dto dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &dto}).GetError(); err != nil {
		return
	}
	userList, total, err := m.Service.GetUserList(&dto)
	if err != nil {
		m.ServerFail(Response{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(Response{
		Data:  userList,
		Total: total,
	})
}

func (m UserApi) UpdateUser(ctx *gin.Context) {
	var dto dto.UserUpdateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &dto}).GetError(); err != nil {
		return
	}
	err := m.Service.UpdateUser(&dto)
	if err != nil {
		m.ServerFail(Response{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(Response{})
}

func (m UserApi) DeleteUser(ctx *gin.Context) {
	var dto dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &dto, BindUri: true}).GetError(); err != nil {
		return
	}
	err := m.Service.DeleteUser(&dto)
	if err != nil {
		m.ServerFail(Response{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(Response{})
}
