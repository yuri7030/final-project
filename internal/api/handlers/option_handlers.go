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
