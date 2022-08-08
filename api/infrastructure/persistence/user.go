package persistence

import (
	"backend/domain/entity"
	"backend/domain/repository"
	"backend/packages/errors"

	"github.com/jinzhu/gorm"
)

type userPersistence struct{}

func NewUserPersistance() repository.UserRepository {
	return &userPersistence{}
}

func (u *userPersistence) Create(db *gorm.DB, user *entity.User) (*entity.User, errors.IError) {
	result := db.Create(&user)

	if result.Error != nil {
		return nil, errors.NewUnexpectedError(result.Error)
	}
	return user, nil
}

func (u *userPersistence) Update(db *gorm.DB, user *entity.User) (*entity.User, errors.IError) {
	result := db.Model(&user).Update(&user)

	if result.Error != nil {
		return nil, errors.NewUnexpectedError(result.Error)
	}
	return user, nil
}

func (u *userPersistence) Delete(db *gorm.DB, id uint) errors.IError {
	user := entity.User{}
	user.ID = id
	if err := db.Delete(&user).Error; err != nil {
		return errors.NewUnexpectedError(err)
	}
	return nil
}

func (u *userPersistence) GetByID(db *gorm.DB, id uint) (*entity.User, errors.IError) {
	user := entity.User{}

	if err := db.Find(&user, id).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.NewNotFoundError(err)
		}
		return nil, errors.NewUnexpectedError(err)
	}
	return &user, nil
}
