package dao

import "github.com/ZhaoJun-hz/go-web-base/model"

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

func (userDao *UserDao) GetUserByNameAndPassword(username string, password string) model.User {
	var user model.User
	tx := userDao.Orm.Model(&user).Where("user_name = ? and password = ?", username, password).Find(&user)
	err := tx.Error
	if err != nil {
		return model.User{}
	}
	return user
}
