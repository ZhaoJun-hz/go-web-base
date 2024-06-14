package dao

import (
	"github.com/ZhaoJun-hz/go-web-base/model"
	"github.com/ZhaoJun-hz/go-web-base/service/dto"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(username string, password string) model.User {
	var user model.User
	tx := userDao.Orm.Model(&user).Where("user_name = ? and password = ?", username, password).Find(&user)
	err := tx.Error
	if err != nil {
		return model.User{}
	}
	return user
}

func (m *UserDao) CheckUserNameExist(username string) bool {
	var total int64
	m.Orm.Model(&model.User{}).Where("user_name = ?", username).Count(&total)
	return total > 0
}

func (m *UserDao) AddUser(userAddDto *dto.UserAddDTO) error {
	var user model.User
	userAddDto.CovertToModel(&user)
	err := userDao.Orm.Save(&user).Error
	if err == nil {
		userAddDto.ID = user.ID
		userAddDto.Password = ""
	}
	return err
}

func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var user model.User
	err := userDao.Orm.First(&user, id).Error
	return user, err
}

func (m *UserDao) GetUserList(dto *dto.UserListDTO) ([]model.User, int64, error) {
	var userList []model.User
	var total int64
	err := m.Orm.Model(&model.User{}).Scopes(Paginate(dto.Paginate)).Find(&userList).
		Offset(-1).Limit(-1).Count(&total).Error
	return userList, total, err
}

func (m *UserDao) UpdateUser(dto *dto.UserUpdateDTO) error {
	var user model.User
	m.Orm.First(&user, dto.ID)
	dto.CovertToModel(&user)
	return m.Orm.Save(&user).Error
}

func (m *UserDao) DeleteUser(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}
