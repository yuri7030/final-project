package entities

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	SurveyId     int
	QuestionText string `gorm:"type:text"`
	AnswerType   int
}

func (Question) TableName() string {
	return "questions"
}
