package api

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
			panic(err)
		}
	}

	database.ConnectDatabase(os.Getenv("DATABASE_URL"))

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{
	// 	"http://localhost:3000",
	// }
	config.AllowHeaders = []string{"Origin", "*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	server.router.Use(cors.New(config))

	routes.InitializeRoutes(server.router)

	return server
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
