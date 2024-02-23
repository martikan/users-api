package service

import (
	"fmt"
	"github.com/martikan/users-api/common"
	"github.com/martikan/users-api/dto"
	"github.com/martikan/users-api/mapper"
	"github.com/martikan/users-api/model"
	"github.com/martikan/users-api/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{repository}
}

func (u UserServiceImpl) CreateUser(dto *dto.CreateUserDTO) error {
	exists, err := u.userRepository.ExistByEmail(dto.Email)
	if err != nil {
		return err
	} else if exists {
		return fmt.Errorf("email is already exists")
	}

	user := &model.User{
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	return u.userRepository.SaveUser(user)
}

func (u UserServiceImpl) GetAllUsers(pageable common.Pageable) ([]dto.UserDTO, error) {
	users, err := u.userRepository.FindUsers(pageable)
	if err != nil {
		return nil, err
	}

	var userDTOs []dto.UserDTO
	for _, u := range users {
		userDTOs = append(userDTOs, mapper.MapUserEntityToDTO(u))
	}

	return userDTOs, nil
}
