package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/config"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/inputs"
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
	exp, err := strconv.Atoi(config.GetValue("JWT_EXPIRY_TIME_SECOND"))

	ttl := time.Duration(exp) * time.Second
	expTime := time.Now().UTC().Add(ttl).Unix()

	token.Claims = jwt.MapClaims{
		"email": login.Email,
		"name":  user.Name,
		"role":  user.Role,
		"id":    user.ID,
		"exp":   expTime,
	}

	tokenString, err := token.SignedString([]byte(config.GetValue("JWT_KEY")))
	if err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to generate JWT token!", nil)
		return
	}

	response := struct {
		Token string `json:"token"`
		Name string `json:"name"`
		Email string `json:"email"`
		Role int `json:"role"`
	}{
		Token: tokenString,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
	}

	common.ResponseSuccess(c, http.StatusOK, "Login successfuly", response)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var register inputs.RegisterInput

	if err := c.ShouldBindJSON(&register); err != nil {
		errorInputs := common.ParseError(err)
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

	if register.Role != 1 && register.Role != 2 {
		common.ResponseError(c, http.StatusBadRequest, "Invalid role", nil)
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

var TokenBlacklist = make(map[string]bool)

func (h *AuthHandler) Logout(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		common.ResponseError(c, http.StatusBadRequest, "Missing Authorization header", nil)
		return
	}

	tokenParts := strings.Split(authorizationHeader, " ")
	token := tokenParts[0]

	TokenBlacklist[token] = true

	common.ResponseSuccess(c, http.StatusOK, "Logged out successfully", nil)
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var input inputs.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}

	user := common.GetUserAuth(c)

	var dbUser entities.User
	if err := database.DB.First(&dbUser, user.ID).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch user data", nil)
		return
	}

	if !common.CheckPasswordHash(input.OldPassword, dbUser.Password) {
		common.ResponseError(c, http.StatusUnauthorized, "Invalid old password", nil)
		return
	}

	newHashPass, err := common.HashPassword(input.NewPassword)
	if err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to hash new password", nil)
		return
	}

	if err := database.DB.Model(&dbUser).Update("password", newHashPass).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to update password", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "Password changed successfully", nil)
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	user := common.GetUserAuth(c)

	var dbUser entities.User
	if err := database.DB.First(&dbUser, user.ID).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch user data", nil)
		return
	}

	response := struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Role int `json:"role"`
	}{
		Name: dbUser.Name,
		Email: dbUser.Email,
		Role: dbUser.Role,
	}

	common.ResponseSuccess(c, http.StatusOK, "Profile fetched successfully", response)
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	var input inputs.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.ResponseError(c, http.StatusBadRequest, "Invalid inputs", common.ParseError(err))
		return
	}

	user := common.GetUserAuth(c)

	var dbUser entities.User
	if err := database.DB.First(&dbUser, user.ID).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch user data", nil)
		return
	}

	dbUser.Name = input.Name
	dbUser.Email = input.Email

	if err := database.DB.Save(&dbUser).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to update profile", nil)
		return
	}

	common.ResponseSuccess(c, http.StatusOK, "Profile updated successfully", nil)
}

func (h* AuthHandler) ValidateToken(c *gin.Context) {
	user := common.GetUserAuth(c)

	var dbUser entities.User
	if err := database.DB.First(&dbUser, user.ID).Error; err != nil {
		common.ResponseError(c, http.StatusInternalServerError, "Failed to fetch user data", nil)
		return
	}

	response := struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Role int `json:"role"`
	}{
		Name: dbUser.Name,
		Email: dbUser.Email,
		Role: dbUser.Role,
	}

	common.ResponseSuccess(c, http.StatusOK, "Token is valid", response)
}