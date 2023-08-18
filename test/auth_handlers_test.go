package handler_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yuri7030/final-project/internal/api/handlers"
)

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTest := NewInitTest()
	initTest.SetUpDb()
	t.Run("Empty email and password", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "",
			"password": "",
		}

		w, _ := initTest.CallApi("/auth/login", payload, nil, handlers.NewAuthHandler().Login, false)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Empty email or password", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "",
			"password": "abc123456",
		}

		w, _ := initTest.CallApi("/auth/login", payload, nil, handlers.NewAuthHandler().Login, false)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Invalid email", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "user_gmail.com",
			"password": "abc123456",
		}

		w, _ := initTest.CallApi("/auth/login", payload, nil, handlers.NewAuthHandler().Login, false)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Wrong email or password", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "user-notfound@gmail.com",
			"password": "abc123456",
		}
		w, _ := initTest.CallApi("/auth/login", payload, nil, handlers.NewAuthHandler().Login, false)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Add success", func(t *testing.T) {
		initTest.CreateDataLogin()
		w, _ := initTest.CallApi("/auth/login", initTest.PayloadLogin, nil, handlers.NewAuthHandler().Login, false)
		log.Println("initTest.PayloadLogin", initTest.PayloadLogin)
		initTest.DeleteDataLogin()

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
