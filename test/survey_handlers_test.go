package handler_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yuri7030/final-project/internal/api/handlers"
	"github.com/yuri7030/final-project/internal/api/inputs"
)

func TestSurveyHandler_CreateSurvey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTest := NewInitTest()
	initTest.SetUpDb()

	handler := handlers.NewSurveyHandler()
	t.Run("Valid input", func(t *testing.T) {
		initTest.CreateDataLogin()
		tokenString := initTest.Login(initTest.PayloadLogin)

		payload := inputs.SurveyCreatingInput{
			Title:       "Test Survey",
			Description: "A test survey",
		}

		headers := map[string]interface{}{
			"Authorization": tokenString,
		}
		w, _ := initTest.CallApi("/surveys", payload, headers, handler.CreateSurvey, true)
		initTest.DeleteDataLogin()

		response := initTest.GetValue(w)
		assert.Equal(t, http.StatusCreated, w.Code)
		data := response["data"].(map[string]interface{})

		assert.Equal(t, "Test Survey", data["title"])
		assert.Equal(t, "A test survey", data["description"])
	})
}
