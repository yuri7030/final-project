package entities

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	QuestionID uint
	OptionText string
	Question   Question `gorm:"foreignKey:QuestionID"`
}

func (Option) TableName() string {
	return "options"
}
