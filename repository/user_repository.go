package repository

import (
	"github.com/martikan/users-api/common"
	"github.com/martikan/users-api/model"
)

type UserRepository interface {
	SaveUser(user *model.User) error
	FindUsers(pageable common.Pageable) ([]model.User, error)
	ExistByEmail(email string) (bool, error)
}
