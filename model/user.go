package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"size:64;not null;unique"`
	RealName string `gorm:"size:64"`
	Avatar   string `gorm:"size:255"`
	Mobile   string `gorm:"size:11"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"size:128;not null" json:"-"`
}
