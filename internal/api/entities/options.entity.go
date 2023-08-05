package entities

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	QuestionId int
	OptionText string
}

func (Option) TableName() string {
	return "options"
}
