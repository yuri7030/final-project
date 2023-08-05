package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/handlers"
	"github.com/yuri7030/final-project/internal/api/middlewares"
)

func InitializeRoutes(router *gin.Engine) {
	authHandler := handlers.NewAuthHandler()
	authGroup := router.Group("/auth")
	authGroup.POST("/login", authHandler.Login)

	surveyHandler := handlers.NewSurveyHandler()
	systemGroup := router.Group("/")
	systemGroup.Use(middlewares.JWTMiddleware())

	systemGroup.POST("surveys", surveyHandler.CreateSurvey)
}
