package service

import (
	"errors"
	"github.com/ZhaoJun-hz/go-web-base/dao"
	"github.com/ZhaoJun-hz/go-web-base/model"
	"github.com/ZhaoJun-hz/go-web-base/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (userService *UserService) Login(userLoginDto dto.UserLoginDTO) (model.User, error) {
	var errResult error
	user := userService.Dao.GetUserByNameAndPassword(userLoginDto.Username, userLoginDto.Password)
	if user.ID == 0 {
		errResult = errors.New("Invalid username or password")
	}
	return user, errResult
}
