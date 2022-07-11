package repository

import (
	"backend/domain/entity"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user *entity.User) (*entity.User, error)
	Update(db *gorm.DB, user *entity.User) (*entity.User, error)
	Delete(db *gorm.DB, id uint) error
	GetByID(db *gorm.DB, id uint) (*entity.User, error)
}
