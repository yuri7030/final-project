package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/constants"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/inputs"
)

type AnswerHandler struct {
}

func NewAnswerHandler() *AnswerHandler {
	return &AnswerHandler{}
}

func (h *AnswerHandler) SubmitSurveyAnswers(c *gin.Context) {
	surveyID, err := strconv.Atoi(c.Param("survey_id"))
	if err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid survey ID", nil)
		return
	}

	var input inputs.SurveyAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}

	var survey entities.Survey
	result := database.DB.Preload("Questions").First(&survey, surveyID)
	if result.RowsAffected == 0 {
		common.ResponseError(c, http.StatusNotFound, "Survey not found", nil)
		return
	}

	guid := uuid.New().String()
	var answers []*entities.Answer
	for _, answerInput := range input.Answers {
		questionID := answerInput.QuestionID
		question := findQuestionByID(survey.Questions, questionID)

		if question == nil {
			continue
		}

		if !validateAnswer(question.AnswerType, answerInput) {
			common.ResponseError(c, http.StatusBadRequest, "Invalid answer format for question", nil)
			return
		}


		answer := entities.Answer{
			QuestionID: question.ID,
			GUID: guid,
			AnswerText: answerInput.AnswerText,
		}

		if question.AnswerType == constants.RadioAnswer {
			answer.SingleOptionID = answerInput.SingleOptionID
		} else if question.AnswerType == constants.CheckboxAnswer {
			answer.MultipleOptionIDs = common.SerializeUintArray(answerInput.MultipleOptionIDs)
		}

		answers = append(answers, &answer)
	}

	if err := database.DB.Create(&answers).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to submit answers", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusCreated, "Answers submitted successfully", answers)
}

func findQuestionByID(questions []*entities.Question, questionID uint) *entities.Question {
	for _, question := range questions {
		if question.ID == questionID {
			return question
		}
	}
	return nil
}

func validateAnswer(answerType int, answerInput inputs.SurveyAnswerInputItem) bool {
	switch answerType {
	case constants.TextAnswer:
		return answerInput.AnswerText != ""
	case constants.RadioAnswer:
		return answerInput.SingleOptionID != 0
	case constants.CheckboxAnswer:
		return answerInput.MultipleOptionIDs != nil && len(answerInput.MultipleOptionIDs) > 0
	default:
		return false
	}
}