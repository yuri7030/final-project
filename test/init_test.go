package handler_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yuri7030/final-project/internal/api/database"
)

func SetUpDb() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		panic(err)
	}

	database.ConnectDatabase()
}

func callApi(apiPath string, handler gin.HandlerFunc, payload interface{}, headers map[string]interface{}) (*httptest.ResponseRecorder, *http.Request) {
	jsonValue, _ := json.Marshal(payload)
	router := gin.Default()
	router.POST(apiPath, handler)
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value.(string))
		}
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w, req
}
