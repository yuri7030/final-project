package entities

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	SurveyId          int
	GuestId           string `gorm:"type:varchar(100)"`
	SingleOptionId    int
	MultipleOptionIds string `gorm:"type:varchar(50)"`
}

func (Answer) TableName() string {
	return "answers"
}
