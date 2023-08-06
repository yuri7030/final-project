package entities

import "gorm.io/gorm"

type Survey struct {
	gorm.Model
	Title     string `gorm:"type:varchar(255)"`
	CreatedBy int
}

func (Survey) TableName() string {
	return "surveys"
}
