package handler_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lithammer/shortuuid"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/database"
	"github.com/yuri7030/final-project/internal/api/entities"
	"github.com/yuri7030/final-project/internal/api/handlers"
	"github.com/yuri7030/final-project/internal/api/middlewares"
)

type InitTest struct {
	PayloadLogin map[string]interface{}
	UserLogin    entities.User
	PlainPass    string
	Email        string
}

type LoginResponseType struct {
	token string
}

func NewInitTest() *InitTest {
	var PlainPass = "abc123456"
	var Email = "user-login-" + shortuuid.New() + "@gmail.com"
	var UserLogin = entities.User{
		Email: Email,
		Role:  1,
		Name:  "user sys",
	}

	var PayloadLogin = map[string]interface{}{
		"email":    UserLogin.Email,
		"password": PlainPass,
	}

	return &InitTest{
		PayloadLogin: PayloadLogin,
		UserLogin:    UserLogin,
		Email:        Email,
		PlainPass:    PlainPass,
	}
}

func (i *InitTest) SetUpDb() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		panic(err)
	}

	database.ConnectDatabase()
}

func (i *InitTest) CallApi(apiPath string, payload interface{}, headers map[string]interface{}, handler gin.HandlerFunc, checkAuth bool) (*httptest.ResponseRecorder, *http.Request) {
	jsonValue, _ := json.Marshal(payload)
	router := gin.Default()
	if checkAuth {
		router.POST(apiPath, middlewares.JWTMiddleware(), handler)
	} else {
		router.POST(apiPath, handler)
	}

	req, _ := http.NewRequest(http.MethodPost, apiPath, bytes.NewBuffer(jsonValue))
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

func (i *InitTest) SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func (i *InitTest) GetValue(res *httptest.ResponseRecorder) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal(res.Body.Bytes(), &result)
	if err != nil {
		log.Println("cannot get value")
	}
	return result
}

func (i *InitTest) Login(userInfo interface{}) string {
	w, _ := i.CallApi("/auth/login", userInfo, nil, handlers.NewAuthHandler().Login, false)
	result := i.GetValue(w)
	return result["token"].(string)
}

func (i *InitTest) CreateDataLogin() bool {
	hashPass, _ := common.HashPassword(i.PlainPass)
	i.UserLogin.Password = hashPass
	database.DB.Create(&i.UserLogin)
	return true
}

func (i *InitTest) DeleteDataLogin() bool {
	database.DB.Unscoped().Delete(&i.UserLogin)
	return true
}
