package handlers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/inputs"
	"github.com/yuri7030/final-project/internal/api/config"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var login inputs.LoginInput

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entities.User
	result := database.DB.Where(&entities.User{Email: login.Email}).First(&user)
	if result.RowsAffected == 0 {
		common.ResponseError(c, http.StatusBadRequest, "This user is not found!", nil)
		return
	}
	
	if !common.CheckPasswordHash(login.Password, user.Password) {
		common.ResponseError(c, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"email": login.Email,
		"name": user.Name,
		"role": user.Role,
		"id": user.ID,
	}

	tokenString, err := token.SignedString([]byte(config.GetValue("JWT_KEY")))
	if err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to generate JWT token!", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var register inputs.RegisterInput

	if err := c.ShouldBindJSON(&register); err != nil {
		errorInputs := common.ParseError(err)
		fmt.Println(errorInputs)
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", errorInputs)
		return
	}

	var oldUser entities.User
	result := database.DB.Where(&entities.User{Email: register.Email}).First(&oldUser)
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
	user.Role = register.Role
	user.Password = hashPass
	database.DB.Create(&user)

	common.ResponseSuccess(c, http.StatusCreated, "Register successfuly", true)
	return
}
