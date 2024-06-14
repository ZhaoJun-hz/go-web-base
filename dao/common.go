package dao

import (
	"github.com/ZhaoJun-hz/go-web-base/service/dto"
	"gorm.io/gorm"
)

func Paginate(p dto.Paginate) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}
