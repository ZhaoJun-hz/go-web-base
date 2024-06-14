package service

import (
	"errors"
	"github.com/ZhaoJun-hz/go-web-base/dao"
	"github.com/ZhaoJun-hz/go-web-base/model"
	"github.com/ZhaoJun-hz/go-web-base/service/dto"
	"github.com/ZhaoJun-hz/go-web-base/utils"
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

func (m *UserService) Login(dto dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string

	user, err := userService.Dao.GetUserByName(dto.Username)
	if err != nil || !utils.CompareHashAndPassword(user.Password, dto.Password) {
		errResult = errors.New("Invalid username or password")
	} else {
		token, err = utils.GeneratorToken(user.ID)
		if err != nil {
			errResult = errors.New("Generate token failed")
		}

	}
	return user, token, errResult
}

func (m *UserService) Register(dto dto.UserRegisterDTO) error {
	if m.Dao.CheckUserNameExist(dto.Username) {
		return errors.New("UserName already exists")
	}
	return m.Dao.Register(dto)
}

func (m *UserService) AddUser(userAddDto *dto.UserAddDTO) error {
	if m.Dao.CheckUserNameExist(userAddDto.UserName) {
		return errors.New("UserName already exists")
	}
	return m.Dao.AddUser(userAddDto)
}

func (m *UserService) GetUserById(commonIdDto *dto.CommonIDDTO) (model.User, error) {
	return m.Dao.GetUserById(commonIdDto.Id)
}

func (m *UserService) GetUserList(dto *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(dto)
}

func (m *UserService) UpdateUser(dto *dto.UserUpdateDTO) error {
	return m.Dao.UpdateUser(dto)
}

func (m *UserService) DeleteUser(dto *dto.CommonIDDTO) error {
	return m.Dao.DeleteUser(dto.Id)
}
