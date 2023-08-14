package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/inputs"	
	"github.com/yuri7030/final-project/internal/constants"

)

type OptionHandler struct {
}

func NewOptionHandler() *OptionHandler {
	return &OptionHandler{}
}

func (h *QuestionHandler) AddOptionsToQuestion(c *gin.Context) {
	var input inputs.OptionsAddingInput
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

	if question.AnswerType != constants.RadioAnswer && 
		question.AnswerType != constants.CheckboxAnswer {
		common.ResponseError(c, http.StatusBadRequest, "Cannot add options for this question type", nil)
		return
	}

	var options []entities.Option
	for _, optionText := range input.OptionTexts {
		option := entities.Option{
			QuestionID: question.ID,
			OptionText: optionText,
		}
		options = append(options, option)
	}

	if err := database.DB.Create(&options).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to add options", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusCreated, "Options added successfully", nil)
}

func (h *QuestionHandler) DeleteOption(c *gin.Context) {
	optionID, err := strconv.Atoi(c.Param("option_id"))
	if err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid option ID", nil)
		return
	}

	var option entities.Option
	result := database.DB.First(&option, optionID)
	if result.RowsAffected == 0 {
		common.ResponseError(c, http.StatusNotFound, "Option not found", nil)
		return
	}

	if err := database.DB.Delete(&option).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to delete option", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "Option deleted successfully", nil)
}

func (h *QuestionHandler) DeleteAllOptions(c *gin.Context) {
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

	if err := database.DB.Where("question_id = ?", question.ID).Delete(&entities.Option{}).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to delete options", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "All options deleted successfully", nil)
}

func (h *QuestionHandler) UpdateOption(c *gin.Context) {
	optionID, err := strconv.Atoi(c.Param("option_id"))
	if err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid option ID", nil)
		return
	}

	var input inputs.OptionUpdatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}

	var option entities.Option
	result := database.DB.First(&option, optionID)
	if result.RowsAffected == 0 {
		common.ResponseError(c, http.StatusNotFound, "Option not found", nil)
		return
	}

	option.OptionText = input.OptionText

	if err := database.DB.Save(&option).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to update option", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "Option updated successfully", nil)
}

func (h *QuestionHandler) ListOptionsByQuestion(c *gin.Context) {
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

	var options []entities.Option
	if err := database.DB.Where("question_id = ?", question.ID).Find(&options).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch options", nil)
		return
	}
	
	var results []map[string]interface{}
	for _, option := range options {
		result := map[string]interface{}{
			"id":          option.ID,
			"optionText":  option.OptionText,
		}
		results = append(results, result)
	}

	common.ResponseSuccess(c, http.StatusOK, "Options fetched successfully", results)
}
