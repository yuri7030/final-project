package inputs

type QuestionCreatingInput struct {
	QuestionText string `json:"questionText" binding:"required"`
	AnswerType   int   `json:"answerType" binding:"required"`
}
