package inputs

type QuestionCreatingInput struct {
	QuestionText string `json:"questionText" binding:"required"`
	AnswerType   uint   `json:"answerType" binding:"required"`
}
