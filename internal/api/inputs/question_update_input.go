package inputs

type QuestionUpdatingInput struct {
	QuestionText string `json:"questionText" binding:"required"`
	AnswerType   uint   `json:"answerType" binding:"required"`
}
