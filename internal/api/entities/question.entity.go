package entities

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	SurveyID     uint
	QuestionText string
	AnswerType   uint

	Survey Survey `gorm:"foreignKey:SurveyID"`
	Answers []Answer `gorm:"foreignKey:QuestionID"`
	Options []Option `gorm:"foreignKey:QuestionID"`
}

// TableName sets the table name for the Question model.
func (Question) TableName() string {
	return "questions"
}
