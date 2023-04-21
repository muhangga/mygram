package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muhangga/config"
	"github.com/muhangga/internal/delivery"
	"github.com/muhangga/internal/repository"
	"github.com/muhangga/internal/service"
	"gorm.io/gorm"
)

type server struct {
	httpServer *gin.Engine
	config     config.Config
}

type Server interface {
	RunServer()
}

func NewServer(config config.Config) Server {
	return &server{
		httpServer: gin.Default(),
		config:     config,
	}
}

func (s *server) DB() *gorm.DB {
	return s.config.Database()
}

func (s *server) RunServer() {

	s.httpServer.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRepository := repository.NewUserRepository(s.DB())

	jwtService := service.NewJwtService()

	authRepository := repository.NewAuthRepository(s.DB())
	authService := service.NewAuthService(authRepository, userRepository)
	authController := delivery.NewAuthDelivery(authService, jwtService)

	api := s.httpServer.Group("/api")

	api.POST("/login", authController.Login)
	api.POST("/register", authController.Register)

	if err := s.httpServer.Run(); err != nil {
		panic(err)
	}
}
