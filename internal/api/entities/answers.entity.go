package entities

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	SurveyID          int
	GUID              string `gorm:"type:varchar(100)"`
	SingleOptionID    int
	MultipleOptionIDs string `gorm:"type:varchar(50)"`
}

func (Answer) TableName() string {
	return "answers"
}
