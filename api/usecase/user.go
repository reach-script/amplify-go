package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/packages/errors"

	"github.com/jinzhu/gorm"
)

type User interface {
	Create(db *gorm.DB, user *entity.User) (*entity.User, errors.IError)
	Update(db *gorm.DB, user *entity.User) (*entity.User, errors.IError)
	Delete(db *gorm.DB, id uint) errors.IError
	GetByID(db *gorm.DB, id uint) (*entity.User, errors.IError)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) User {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Create(db *gorm.DB, user *entity.User) (*entity.User, errors.IError) {
	return u.userRepository.Create(db, user)
}

func (u *userUseCase) Update(db *gorm.DB, user *entity.User) (*entity.User, errors.IError) {
	return u.userRepository.Update(db, user)
}

func (u *userUseCase) Delete(db *gorm.DB, id uint) errors.IError {
	return u.userRepository.Delete(db, id)
}

func (u *userUseCase) GetByID(db *gorm.DB, id uint) (*entity.User, errors.IError) {
	return u.userRepository.GetByID(db, id)
}
