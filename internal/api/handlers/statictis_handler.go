package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
)

type StatictisHandler struct {
}

func NewStatictisHandler() *StatictisHandler {
	return &StatictisHandler{}
}

func (h *StatictisHandler) GetSurveyWithMostRespondents(c *gin.Context) {
	user := common.GetUserAuth(c)

	type Result struct {
		SurveyID        uint   `json:"survey_id"`
		SurveyTitle     string `json:"survey_title"`
		RespondentCount int    `json:"respondent_count"`
	}

	var result *Result

	selectField := []string{
		"surveys.id as SurveyID",
		"surveys.title as SurveyTitle",
		"COUNT(DISTINCT answers.guid) AS RespondentCount",
	}

	database.DB.Table("surveys").Select(selectField).Joins("join questions ON questions.survey_id = surveys.id").Joins("join answers ON answers.question_id = questions.id").Where("surveys.created_by = ?", user.ID).Group("surveys.id").Order("RespondentCount DESC").Limit(1).Scan(&result)

	common.ResponseSuccess(c, http.StatusOK, "Survey with most respondents retrieved successfully", result)
	return
}

func (h *StatictisHandler) GetSurveyWithLeastRespondents(c *gin.Context) {
	user := common.GetUserAuth(c)

	type Result struct {
		SurveyID        uint   `json:"survey_id"`
		SurveyTitle     string `json:"survey_title"`
		RespondentCount int    `json:"respondent_count"`
	}

	var result *Result

	selectField := []string{
		"surveys.id as SurveyID",
		"surveys.title as SurveyTitle",
		"COUNT(DISTINCT answers.guid) AS RespondentCount",
	}

	database.DB.Table("surveys").Select(selectField).Joins("join questions ON questions.survey_id = surveys.id").Joins("join answers ON answers.question_id = questions.id").Where("surveys.created_by = ?", user.ID).Group("surveys.id").Order("RespondentCount ASC").Limit(1).Scan(&result)

	common.ResponseSuccess(c, http.StatusOK, "Survey with most respondents retrieved successfully", result)
	return
}
