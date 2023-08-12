package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/handlers"
	"github.com/yuri7030/final-project/internal/api/middlewares"
)

func InitializeRoutes(router *gin.Engine) {
	router.Use(gin.Recovery())

	authHandler := handlers.NewAuthHandler()
	authGroup := router.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/logout", authHandler.Logout)
	surveyHandler := handlers.NewSurveyHandler()
	systemGroup := router.Group("/")
	systemGroup.Use(middlewares.JWTMiddleware())

	systemGroup.POST("/surveys", surveyHandler.CreateSurvey)
}
