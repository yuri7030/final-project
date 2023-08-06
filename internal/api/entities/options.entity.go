package entities

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	QuestionID int
	OptionText string
}

func (Option) TableName() string {
	return "options"
}
