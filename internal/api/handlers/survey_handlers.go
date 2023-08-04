package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/yuri7030/final-project/internal/api/models"
)

type SurveyHandler struct {
}

var surveys []models.Survey

func NewSurveyHandler() *SurveyHandler {
	return &SurveyHandler{}
}

func (h *SurveyHandler) CreateSurvey(c *gin.Context) {
	var survey models.Survey
	if err := c.ShouldBindJSON(&survey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	survey.ID = len(surveys) + 1
	surveys = append(surveys, survey)
	c.JSON(http.StatusCreated, survey)
}