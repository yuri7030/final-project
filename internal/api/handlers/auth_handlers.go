package handlers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/constants"
)

type AuthHandler struct {
}

type registerInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(constants.JwtSecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var register registerInput

	if err := c.ShouldBindJSON(&register); err != nil {
		errorInputs := common.ParseError(err)
		fmt.Println(errorInputs)
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", errorInputs)
		return
	}

	var oldUser entities.User
	result := database.DB.Debug().Where(&entities.User{Email: register.Email}).First(oldUser)
	if result.RowsAffected > 0 {
		common.ResponseError(c, http.StatusBadRequest, "Email already exists", nil)
		return
	}

	hashPass, err := common.HashPassword(register.Password)
	if err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Something went wrong", nil)
		return
	}
	var user entities.User
	user.Name = register.Name
	user.Email = register.Email
	user.Password = hashPass
	database.DB.Create(&user)

	common.ResponseSuccess(c, http.StatusCreated, "Register successfuly", true)
	return
}
