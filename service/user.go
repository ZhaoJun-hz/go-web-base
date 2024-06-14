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

func (m *UserService) Login(userLoginDto dto.UserLoginDTO) (model.User, error) {
	var errResult error
	user := userService.Dao.GetUserByNameAndPassword(userLoginDto.Username, userLoginDto.Password)
	if user.ID == 0 {
		errResult = errors.New("Invalid username or password")
	}
	return user, errResult
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
