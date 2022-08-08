package repository

import (
	"backend/domain/entity"
	"backend/packages/errors"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user *entity.User) (*entity.User, errors.IError)
	Update(db *gorm.DB, user *entity.User) (*entity.User, errors.IError)
	Delete(db *gorm.DB, id uint) errors.IError
	GetByID(db *gorm.DB, id uint) (*entity.User, errors.IError)
}
