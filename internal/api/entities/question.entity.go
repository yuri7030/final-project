package entities

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	SurveyID     int
	QuestionText string `gorm:"type:text"`
	AnswerType   int
}

func (Question) TableName() string {
	return "questions"
}
