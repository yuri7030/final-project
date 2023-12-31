package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/inputs"
	"github.com/yuri7030/final-project/internal/constants/answer_type_enums"
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

func (h *QuestionHandler) AddMultipleQuestionToSurvey(c *gin.Context) {
	var input inputs.QuestionMultiCreatingInput
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

	result := make([]interface{}, 0)
	for _, q := range input.Questions {
		question := entities.Question{
			SurveyID:     survey.ID,
			QuestionText: q.QuestionText,
			AnswerType:   q.AnswerType,
		}
		if err := database.DB.Create(&question).Error; err != nil {
			common.ResponseError(c, http.StatusInternalServerError, "Failed to create question", nil)
			return
		}
		curResult := map[string]interface{}{
			"id":           question.ID,
			"questionText": question.QuestionText,
			"answerType":   question.AnswerType,
		}
		result = append(result, curResult)
	}

	common.ResponseSuccess(c, http.StatusCreated, "Questions added successfully", result)
}

func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	var input inputs.QuestionUpdatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}

	questionID, err := strconv.Atoi(c.Param("question_id"))
	if err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid question ID", nil)
		return
	}

	var question entities.Question
	result := database.DB.First(&question, questionID)
	if result.RowsAffected == 0 {
		common.ResponseError(c, http.StatusNotFound, "Question not found", nil)
		return
	}

	question.QuestionText = input.QuestionText
	question.AnswerType = input.AnswerType

	if err := database.DB.Save(&question).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to update question", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "Question updated successfully", nil)
}

func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	questionID, err := strconv.Atoi(c.Param("question_id"))
	if err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid question ID", nil)
		return
	}

	var question entities.Question
	result := database.DB.First(&question, questionID)
	if result.RowsAffected == 0 {
		common.ResponseError(c, http.StatusNotFound, "Question not found", nil)
		return
	}

	if err := database.DB.Delete(&question).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to delete question", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "Question deleted successfully", nil)
}

func (h *QuestionHandler) ListQuestionsBySurvey(c *gin.Context) {
	surveyID, err := strconv.Atoi(c.Param("survey_id"))
	if err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid survey ID", nil)
		return
	}

	var survey entities.Survey
	result := database.DB.First(&survey, surveyID)
	if result.RowsAffected == 0 {
		common.ResponseError(c, http.StatusNotFound, "Survey not found", nil)
		return
	}

	var questions []entities.Question
	if err := database.DB.Where("survey_id = ?", survey.ID).Find(&questions).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch questions", nil)
		return
	}

	var results []map[string]interface{}
	for _, question := range questions {
		result := map[string]interface{}{
			"id":           question.ID,
			"questionText": question.QuestionText,
			"answerType":   answer_type_enums.AnswerTypes[question.AnswerType],
		}
		results = append(results, result)
	}

	common.ResponseSuccess(c, http.StatusOK, "Questions listed successfully", results)
}
