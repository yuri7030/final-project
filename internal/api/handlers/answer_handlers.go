package handlers

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/constants/answer_type_enums"
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

		if question.AnswerType == answer_type_enums.RadioAnswer {
			answer.SingleOptionID = answerInput.SingleOptionID
		} else if question.AnswerType == answer_type_enums.CheckboxAnswer {
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
	case answer_type_enums.TextAnswer:
		return answerInput.AnswerText != ""
	case answer_type_enums.RadioAnswer:
		return answerInput.SingleOptionID != 0
	case answer_type_enums.CheckboxAnswer:
		return answerInput.MultipleOptionIDs != nil && len(answerInput.MultipleOptionIDs) > 0
	default:
		return false
	}
}

func (h *AnswerHandler) AggregateSurveyAnswers(c *gin.Context) {
    surveyID, err := strconv.Atoi(c.Param("survey_id"))
	if err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid survey ID", nil)
		return
	}

	var answers []*entities.Answer
	result := database.DB.Where("question_id IN (SELECT id FROM questions WHERE survey_id = ?)", surveyID).Find(&answers)
	if result.Error != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch answers", nil)
		return
	}

	var uniqueGUIDs []string
	countByGUID := make(map[string]int)

	for _, answer := range answers {
		if _, exists := countByGUID[answer.GUID]; !exists {
			uniqueGUIDs = append(uniqueGUIDs, answer.GUID)
		}
		countByGUID[answer.GUID]++
	}

	response := struct {
		TotalRespondents int        `json:"total_respondents"`
	}{
		TotalRespondents: len(uniqueGUIDs),
	}

	common.ResponseSuccess(c, http.StatusOK, "Survey response count retrieved successfully", response)
}