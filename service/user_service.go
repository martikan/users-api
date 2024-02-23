package service

import (
	"github.com/martikan/users-api/common"
	"github.com/martikan/users-api/dto"
)

type UserService interface {
	CreateUser(dto *dto.CreateUserDTO) error
	GetAllUsers(pageable common.Pageable) ([]dto.UserDTO, error)
}
