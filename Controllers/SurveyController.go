package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnswerType string

const (
	AnswerTypeNumber         AnswerType = "number"
	AnswerTypeString         AnswerType = "string"
	AnswerTypeDropdown       AnswerType = "dropdown"
	AnswerTypeMultipleChoice AnswerType = "multiple_choice"
	// Add more answer types as needed
)

type Question struct {
	QuestionID    int        `json:"question_id"`
	Questionnaire string     `json:"questionnaire"`
	AnswerType    AnswerType `json:"answer_type"`
}

type Survey struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

var surveys []Survey

func CreateSurvey(c *gin.Context) {
	var survey Survey
	if err := c.ShouldBindJSON(&survey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	surveys = append(surveys, survey)
	c.JSON(http.StatusCreated, survey)
}

func GetQuestions(c *gin.Context) {
	surveyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid survey ID"})
		return
	}

	for _, survey := range surveys {
		if survey.ID == surveyID {
			c.JSON(http.StatusOK, survey.Questions)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Survey not found"})
}

func UpdateQuestions(c *gin.Context) {
	surveyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid survey ID"})
		return
	}

	var updateSurvey Survey
	if err := c.ShouldBindJSON(&updateSurvey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, survey := range surveys {
		if survey.ID == surveyID {
			surveys[i].Questions = updateSurvey.Questions
			c.JSON(http.StatusOK, surveys[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Survey not found"})
}
