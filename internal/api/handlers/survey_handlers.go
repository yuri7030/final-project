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

    user, exists := c.Get("user")
    if !exists {
        common.ResponseError(c, http.StatusUnauthorized, "Unauthorized", nil)
        return
    }

    currentUser := user.(entities.User)

    survey := entities.Survey{
        Title:       input.Title,
        Description: input.Description,
        CreatedBy: currentUser.ID,
    }

    if err := database.DB.Create(&survey).Error; err != nil {
        common.ResponseError(c, http.StatusInternalServerError, "Failed to create survey", nil)
        return
    }

    result := map[string]interface{} {
        "id":          survey.ID,
        "title":       survey.Title,
        "description": survey.Description,
    }

    common.ResponseSuccess(c, http.StatusCreated, "Survey created successfully", result)
}
