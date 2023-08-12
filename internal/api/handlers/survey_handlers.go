package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/inputs"
)

type SurveyHandler struct {
}

func NewSurveyHandler() *SurveyHandler {
	return &SurveyHandler{}
}

func (h *SurveyHandler) CreateSurvey(c *gin.Context) {
	var input inputs.SurveyCreatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}

	user := common.GetUserAuth(c)

	survey := entities.Survey{
		Title:       input.Title,
		Description: input.Description,
		CreatedBy:   user.ID,
	}

	if err := database.DB.Create(&survey).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to create survey", nil)
		return
	}

	result := map[string]interface{}{
		"id":          survey.ID,
		"title":       survey.Title,
		"description": survey.Description,
	}

	common.ResponseSuccess(c, http.StatusCreated, "Survey created successfully", result)
}

func (h *SurveyHandler) UpdateSurvey(c *gin.Context) {
	var input inputs.SurveyUpdatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}
	surveyID := c.Param("id")

	var survey entities.Survey
	if err := database.DB.First(&survey, surveyID).Error; err != nil {
		common.ResponseError(c, http.StatusNotFound, "Survey not found", nil)
		return
	}

	survey.Title = input.Title
	survey.Description = input.Description

	if err := database.DB.Save(&survey).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to update survey", nil)
		return
	}

	result := map[string]interface{}{
		"id":          survey.ID,
		"title":       survey.Title,
		"description": survey.Description,
	}

	common.ResponseSuccess(c, http.StatusOK, "Survey updated successfully", result)
}

func (h *SurveyHandler) DeleteSurvey(c *gin.Context) {
	surveyID := c.Param("id")

	var survey entities.Survey
	if err := database.DB.First(&survey, surveyID).Error; err != nil {
		common.ResponseError(c, http.StatusNotFound, "Survey not found", nil)
		return
	}

	if err := database.DB.Delete(&survey).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to delete survey", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "Survey deleted successfully", nil)
}

func (h *SurveyHandler) ListSurveysByCurrentUser(c *gin.Context) {
	user := common.GetUserAuth(c)
	var surveys []entities.Survey
	if err := database.DB.Where("created_by = ?", user.ID).Find(&surveys).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch surveys", nil)
		return
	}
	var results []map[string]interface{}
	for _, survey := range surveys {
		result := map[string]interface{}{
			"id":          survey.ID,
			"title":       survey.Title,
			"description": survey.Description,
		}
		results = append(results, result)
	}

	common.ResponseSuccess(c, http.StatusOK, "Surveys fetched successfully", results)
}
