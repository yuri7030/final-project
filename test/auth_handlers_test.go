package handler_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/assert"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/handlers"
)

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	SetUpDb()
	t.Run("Empty email and password", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "",
			"password": "",
		}

		w, _ := callApi("/auth/login", handlers.NewAuthHandler().Login, payload, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Empty email or password", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "",
			"password": "abc123456",
		}

		w, _ := callApi("/auth/login", handlers.NewAuthHandler().Login, payload, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Invalid email", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "user_gmail.com",
			"password": "abc123456",
		}

		w, _ := callApi("/auth/login", handlers.NewAuthHandler().Login, payload, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Wrong email or password", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":    "user-notfound@gmail.com",
			"password": "abc123456",
		}
		w, _ := callApi("/auth/login", handlers.NewAuthHandler().Login, payload, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	sUUID := shortuuid.New()
	plainpass := "abc123456"
	email := "user-login-" + sUUID + "@gmail.com"

	userLogin := entities.User{
		Email: email,
		Role:  1,
		Name:  "user sys",
	}

	createDataLogin := func() bool {
		hashPass, _ := common.HashPassword(plainpass)
		userLogin.Password = hashPass
		database.DB.Create(&userLogin)
		return true
	}

	deleteDataLogin := func() bool {
		database.DB.Unscoped().Delete(&userLogin)
		return true
	}

	t.Run("Add success", func(t *testing.T) {
		createDataLogin()
		payload := map[string]interface{}{
			"email":    userLogin.Email,
			"password": plainpass,
		}

		w, _ := callApi("/auth/login", handlers.NewAuthHandler().Login, payload, nil)
		deleteDataLogin()

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
