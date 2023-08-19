package inputs

type SurveyAnswerInput struct {
	Answers []SurveyAnswerInputItem `json:"answers" binding:"required"`
}

type SurveyAnswerInputItem struct {
	QuestionID       uint   `json:"questionId" binding:"required"`
	AnswerText       string `json:"answerText"`
	SingleOptionID   uint   `json:"singleOptionId"`
	MultipleOptionIDs []uint `json:"multipleOptionIds"`
}
