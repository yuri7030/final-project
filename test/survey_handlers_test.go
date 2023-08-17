package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/handlers"
	"github.com/yuri7030/final-project/internal/api/inputs"
)

func TestSurveyHandler_CreateSurvey(t *testing.T) {
	router := SetupTestRouter()
	database.ConnectDatabase()

	handler := handlers.NewSurveyHandler()

	t.Run("Valid input", func(t *testing.T) {
		input := inputs.SurveyCreatingInput{
			Title:       "Test Survey",
			Description: "A test survey",
		}
		jsonValue, _ := json.Marshal(input)

		req, _ := http.NewRequest(http.MethodPost, "/surveys", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		router.POST("/surveys", handler.CreateSurvey)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Test Survey", response["title"])
		assert.Equal(t, "A test survey", response["description"])
	})
}
