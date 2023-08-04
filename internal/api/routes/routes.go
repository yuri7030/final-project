package routes

import (
	"github.com/yuri7030/final-project/internal/api/handlers"
	"github.com/yuri7030/final-project/internal/api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.Use(middlewares.AuthMiddleware)

	authGroup := router.Group("/auth")
	authHandler := handlers.NewAuthHandler()
	authGroup.POST("/login", authHandler.Login)
}