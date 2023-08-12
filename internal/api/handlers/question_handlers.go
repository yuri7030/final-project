package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/inputs"
)

type QuestionHandler struct {
}

func NewQuestionHandler() *QuestionHandler {
	return &QuestionHandler{}
}

func (h *QuestionHandler) AddQuestionToSurvey(c *gin.Context) {
	var input inputs.QuestionCreatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}

	surveyID := c.Param("survey_id")
	var survey entities.Survey
	if err := database.DB.First(&survey, surveyID).Error; err != nil {
		common.ResponseError(c, http.StatusNotFound, "Survey not found", nil)
		return
	}

	question := entities.Question{
		SurveyID:     survey.ID,
		QuestionText: input.QuestionText,
		AnswerType:   input.AnswerType,
	}

	if err := database.DB.Create(&question).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to create question", nil)
		return
	}

	result := map[string]interface{}{
		"id":           question.ID,
		"questionText": question.QuestionText,
		"answerType":   question.AnswerType,
	}

	common.ResponseSuccess(c, http.StatusCreated, "Question added successfully", result)
}
