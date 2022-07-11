package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"github.com/jinzhu/gorm"
)

type User interface {
	Create(db *gorm.DB, user *entity.User) (*entity.User, error)
	Update(db *gorm.DB, user *entity.User) (*entity.User, error)
	Delete(db *gorm.DB, id uint) error
	GetByID(db *gorm.DB, id uint) (*entity.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) User {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Create(db *gorm.DB, user *entity.User) (*entity.User, error) {
	return u.userRepository.Create(db, user)
}

func (u *userUseCase) Update(db *gorm.DB, user *entity.User) (*entity.User, error) {
	return u.userRepository.Update(db, user)
}

func (u *userUseCase) Delete(db *gorm.DB, id uint) error {
	return u.userRepository.Delete(db, id)
}

func (u *userUseCase) GetByID(db *gorm.DB, id uint) (*entity.User, error) {
	return u.userRepository.GetByID(db, id)
}
