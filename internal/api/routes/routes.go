package routes

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/handlers"
	"github.com/yuri7030/final-project/internal/api/middlewares"
)

func InitializeRoutes(router *gin.Engine) {
	router.Use(gin.Recovery())


	router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "ok",
        })
    })

	authHandler := handlers.NewAuthHandler()
	authGroup := router.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/register", authHandler.Register)

	surveyHandler := handlers.NewSurveyHandler()
	questionHandler := handlers.NewQuestionHandler()
	answerHandler := handlers.NewAnswerHandler()
	statisticsHandler := handlers.NewStatictisHandler()

	backOfficeGroup := router.Group("/")
	backOfficeGroup.Use(middlewares.JWTMiddleware())
	backOfficeGroup.Use(middlewares.BlacklistMiddleware())

	backOfficeGroup.POST("/logout", authHandler.Logout)
	backOfficeGroup.POST("/change-password", authHandler.ChangePassword)
	backOfficeGroup.GET("/me", authHandler.GetProfile)
	backOfficeGroup.PUT("/me", authHandler.UpdateProfile)
	backOfficeGroup.POST("/validate-token", authHandler.ValidateToken)

	surveyGroup := backOfficeGroup.Group("/surveys")
	{
		surveyGroup.GET("", surveyHandler.ListSurveysByCurrentUser)
		surveyGroup.POST("", surveyHandler.CreateSurvey)
		surveyGroup.PUT("/:survey_id", surveyHandler.UpdateSurvey)
		surveyGroup.DELETE("/:survey_id", surveyHandler.DeleteSurvey)
		surveyGroup.GET("/:survey_id/questions", questionHandler.ListQuestionsBySurvey)
		surveyGroup.POST("/:survey_id/questions", questionHandler.AddQuestionToSurvey)
		surveyGroup.POST("/:survey_id/multi-questions/", questionHandler.AddMultipleQuestionToSurvey)
		surveyGroup.POST("/:survey_id/submit", answerHandler.SubmitSurveyAnswers)
		surveyGroup.GET("/:survey_id/number-people-answer", answerHandler.AggregateSurveyAnswers)
	}

	questionGroup := backOfficeGroup.Group("/questions")
	{
		questionGroup.PUT("/:question_id", questionHandler.UpdateQuestion)
		questionGroup.DELETE("/:question_id", questionHandler.DeleteQuestion)
		questionGroup.POST("/:question_id/options", questionHandler.AddOptionsToQuestion)
		questionGroup.DELETE("/:question_id/options", questionHandler.DeleteAllOptions)
		questionGroup.GET("/:question_id/options", questionHandler.ListOptionsByQuestion)
	}

	optionsGroup := backOfficeGroup.Group("/options")
	{
		optionsGroup.DELETE("/:option_id", questionHandler.DeleteOption)
		optionsGroup.PUT("/:option_id", questionHandler.UpdateOption)
	}

	statisticGroup := backOfficeGroup.Group("/statistics")
	{
		statisticGroup.GET("/surveys/most-respondents", statisticsHandler.GetSurveyWithMostRespondents)
		statisticGroup.GET("/surveys/least-respondents", statisticsHandler.GetSurveyWithLeastRespondents)
	}
}
