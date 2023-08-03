package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuri7030/final-project/Controllers"
)

func main() {
	r := gin.Default()

	// Define the endpoints
	r.POST("/surveys", Controllers.CreateSurvey)
	r.GET("/surveys/:id/questions", Controllers.GetQuestions)
	r.PUT("/surveys/:id/questions", Controllers.UpdateQuestions)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
