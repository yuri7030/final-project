package entities

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	QuestionID 		  uint
	GUID              string `gorm:"type:varchar(50)"`
	AnswerText        string
	SingleOptionID    uint
	MultipleOptionIDs string `gorm:"type:varchar(50)"`
	
	Question Question `gorm:"foreignKey:QuestionID"`
}

func (Answer) TableName() string {
	return "answers"
}
