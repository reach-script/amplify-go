package persistence

import (
	"backend/domain/entity"
	"backend/domain/repository"

	"github.com/jinzhu/gorm"
)

type userPersistence struct{}

func NewUserPersistance() repository.UserRepository {
	return &userPersistence{}
}

func (u *userPersistence) Create(db *gorm.DB, user *entity.User) (*entity.User, error) {
	result := db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *userPersistence) Update(db *gorm.DB, user *entity.User) (*entity.User, error) {
	result := db.Model(&user).Update(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *userPersistence) Delete(db *gorm.DB, id uint) error {
	user := entity.User{}
	user.ID = id
	return db.Delete(&user).Error
}

func (u *userPersistence) GetByID(db *gorm.DB, id uint) (*entity.User, error) {
	user := entity.User{}
	result := db.Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
