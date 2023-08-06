package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/config"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/routes"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}

	_, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	database.ConnectDatabase()

	routes.InitializeRoutes(server.router)

	return server
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
