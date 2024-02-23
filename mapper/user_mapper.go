package mapper

import (
	"github.com/martikan/users-api/dto"
	"github.com/martikan/users-api/model"
)

func MapUserEntityToDTO(e model.User) dto.UserDTO {
	return dto.UserDTO{
		ID:        e.ID,
		Email:     e.Email,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
}
