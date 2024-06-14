package model

import (
	"github.com/ZhaoJun-hz/go-web-base/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"size:64;not null;unique"`
	RealName string `gorm:"size:64"`
	Avatar   string `gorm:"size:255"`
	Mobile   string `gorm:"size:11"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"size:128;not null" json:"-"`
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	return m.Encrypt()
}

func (m *User) Encrypt() error {
	encrypt, err := utils.Encrypt(m.Password)
	if err == nil {
		m.Password = encrypt
	}
	return err
}
