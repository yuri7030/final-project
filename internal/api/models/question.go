package models

type AnswerType string

type Question struct {
	QuestionID  int       `json:"question_id"`
	Questionnaire string   `json:"questionnaire"`
	AnswerType  AnswerType `json:"answer_type"`
}