package handlers

type QuestionHandler struct {
}

func NewQuestionHandler() *QuestionHandler {
	return &QuestionHandler{}
}

// func GetQuestions(c *gin.Context) {
// 	surveyID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid survey ID"})
// 		return
// 	}

// 	for _, survey := range surveyList {
// 		if survey.ID == surveyID {
// 			c.JSON(http.StatusOK, survey.Questions)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "Survey not found"})
// }

// func UpdateQuestions(c *gin.Context) {
// 	surveyID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid survey ID"})
// 		return
// 	}

// 	var updateSurvey models.Survey
// 	if err := c.ShouldBindJSON(&updateSurvey); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for i, survey := range surveyList {
// 		if survey.ID == surveyID {
// 			surveys[i].Questions = updateSurvey.Questions
// 			c.JSON(http.StatusOK, surveys[i])
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "Survey not found"})
// }
