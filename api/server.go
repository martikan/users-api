package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/martikan/users-api/repository"
	"github.com/martikan/users-api/service"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(database *gorm.DB) (*Server, error) {
	server := &Server{
		db: database,
	}
	server.SetupRouter()
	return server, nil
}

func (s *Server) SetupRouter() {
	router := gin.Default()

	basePath := "/api"

	v1Grp := router.Group(fmt.Sprintf("%s/v1", basePath))

	// User handler
	userHandler := s.initUserHandler()
	v1Grp.GET("/users", userHandler.getUsers)
	v1Grp.POST("/users", userHandler.createUser)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s *Server) initUserHandler() *UserController {
	userRepository := repository.NewUserRepositoryImpl(s.db)
	userService := service.NewUserServiceImpl(userRepository)
	return s.NewUserController(userService)
}
