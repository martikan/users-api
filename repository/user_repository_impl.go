package repository

import (
	"github.com/martikan/users-api/common"
	"github.com/martikan/users-api/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (u UserRepositoryImpl) SaveUser(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepositoryImpl) FindUsers(pageable common.Pageable) ([]model.User, error) {
	var users []model.User
	if err := u.db.Limit(pageable.GetSize()).Offset(pageable.GetOffset()).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepositoryImpl) ExistByEmail(email string) (bool, error) {
	var exists int64
	if err := u.db.Model(&model.User{}).Where("email = ?", email).Count(&exists).Error; err != nil {
		return false, err
	}
	return exists > 0, nil
}
