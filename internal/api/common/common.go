package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yuri7030/final-project/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

func ResponseSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  true,
		"data":    data,
		"message": message,
	})
	return
}

func ResponseError(c *gin.Context, statusCode int, message string, errs interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  false,
		"message": message,
		"errors":  errs,
	})
	return
}

func ParseError(err error) []string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {
			var message string
			tag := e.Tag()

			if _, exists := validations[tag]; exists {
				message = validations[tag]
			}

			switch tag {
			case "required", "email":
				message = fmt.Sprintf(message, e.Field())
			case "required_special":
				message = fmt.Sprintf("The field %s is required if %s is not supplied", e.Field(), e.Param())
			default:
				message = fmt.Sprintf(validations["_default"], e.Field())
			}

			errorMessages[i] = message
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{fmt.Sprintf("The field %s must be a %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), constants.PasswordHashLength)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserAuth(c *gin.Context) AuthJWT {
	user, err := c.Get("user")
	if !err {
		panic("User auth not exist")
	}
	return user.(AuthJWT)
}

func PrettyPrint(v interface{}) {
	str, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println(string(str))
}

func SerializeUintArray(arr []uint) string {
	strArr := make([]string, len(arr))
	for i, val := range arr {
		strArr[i] = strconv.FormatUint(uint64(val), 10)
	}
	return strings.Join(strArr, ",")
}
