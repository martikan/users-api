package api

import (
	"github.com/gin-gonic/gin"
	"github.com/martikan/users-api/common"
	"github.com/martikan/users-api/dto"
	"github.com/martikan/users-api/service"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func (s *Server) NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) getUsers(ctx *gin.Context) {
	var pageable common.DefaultPageable
	if err := ctx.ShouldBindQuery(&pageable); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	users, err := u.userService.GetAllUsers(pageable)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) createUser(ctx *gin.Context) {
	var req dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := u.userService.CreateUser(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusCreated)
}
