package entities

import "gorm.io/gorm"

type Survey struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255)"`
	Description string
	CreatedBy   uint

	Creator   *User       `gorm:"foreignKey:CreatedBy"`
	Questions []*Question `gorm:"foreignKey:SurveyID"`
}

func (Survey) TableName() string {
	return "surveys"
}
