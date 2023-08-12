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
	authGroup.POST("/logout", middlewares.JWTMiddleware(), middlewares.BlacklistMiddleware(), authHandler.Logout)
	authGroup.POST("/change-password", middlewares.JWTMiddleware(), middlewares.BlacklistMiddleware(), authHandler.ChangePassword)

	surveyHandler := handlers.NewSurveyHandler()
	surveyGroup := router.Group("/survey")
	surveyGroup.Use(middlewares.JWTMiddleware())
	surveyGroup.Use(middlewares.BlacklistMiddleware())
	surveyGroup.POST("/", surveyHandler.CreateSurvey)
	surveyGroup.PUT("/:id", surveyHandler.UpdateSurvey)
	surveyGroup.DELETE("/:id", surveyHandler.DeleteSurvey)
	surveyGroup.GET("/surveys/my", surveyHandler.ListSurveysByCurrentUser)

	questionHandler := handlers.NewQuestionHandler()
	questionGroup := router.Group("/survey/question")
	questionGroup.Use(middlewares.JWTMiddleware())
	questionGroup.Use(middlewares.BlacklistMiddleware())
	questionGroup.POST("/:survey_id", questionHandler.AddQuestionToSurvey)
	questionGroup.PUT("/:question_id", questionHandler.UpdateQuestion)
	questionGroup.DELETE("/:question_id", questionHandler.DeleteQuestion)
}
