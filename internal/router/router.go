package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muhangga/config"
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

	if err := s.httpServer.Run(); err != nil {
		panic(err)
	}
}
