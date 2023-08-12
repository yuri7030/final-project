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
	surveyGroup := router.Group("/survey")
	surveyGroup.Use(middlewares.JWTMiddleware())
	surveyGroup.POST("/", surveyHandler.CreateSurvey)
	surveyGroup.PUT("/:id", surveyHandler.UpdateSurvey)
	surveyGroup.DELETE("/:id", surveyHandler.DeleteSurvey)
	surveyGroup.GET("/surveys/my", surveyHandler.ListSurveysByCurrentUser)

	questionHandler := handlers.NewQuestionHandler()
	questionGroup := router.Group("/survey/:survey_id/question")
	questionGroup.Use(middlewares.JWTMiddleware())
	questionGroup.POST("/", questionHandler.AddQuestionToSurvey)
}
