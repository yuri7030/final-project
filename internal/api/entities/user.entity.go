package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(255)"`
	Role     int    // 1: admin, 2: user

	Surveys []*Survey `gorm:"foreignKey:CreatedBy"`
}

func (User) TableName() string {
	return "users"
}
