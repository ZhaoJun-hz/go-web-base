package dto

import (
	"github.com/ZhaoJun-hz/go-web-base/model"
)

type UserLoginDTO struct {
	Username string `json:"username" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

type UserRegisterDTO struct {
	Username string `json:"username" binding:"required" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

func (m *UserRegisterDTO) CovertToModel(user *model.User) {
	user.UserName = m.Username
	user.Password = m.Password
}

type UserAddDTO struct {
	ID       uint
	UserName string `json:"user_name" form:"user_name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" gorm:"password" binding:"required" message:"密码不能为空"`
}

func (m *UserAddDTO) CovertToModel(user *model.User) {
	user.UserName = m.UserName
	user.RealName = m.RealName
	user.Mobile = m.Mobile
	user.Password = m.Password
	user.Avatar = m.Avatar
}

type UserListDTO struct {
	Paginate
}

type UserUpdateDTO struct {
	ID       uint   `json:"id" binding:"required" message:"id不能为空"`
	UserName string `json:"user_name" gorm:"user_name"`
	RealName string `json:"real_name" gorm:"real_name"`
	Mobile   string `json:"mobile" gorm:"mobile"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

func (m *UserUpdateDTO) CovertToModel(user *model.User) {
	user.ID = m.ID
	user.UserName = m.UserName
	user.RealName = m.RealName
	user.Mobile = m.Mobile
	user.Password = m.Password
}
