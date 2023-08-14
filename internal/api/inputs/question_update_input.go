package inputs

type QuestionUpdatingInput struct {
	QuestionText string `json:"questionText" binding:"required"`
	AnswerType   int   `json:"answerType" binding:"required"`
}
