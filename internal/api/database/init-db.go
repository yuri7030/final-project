package database

import (
	"github.com/yuri7030/final-project/internal/api/entities"
  	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(connString string) {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(
		&entities.User{},
		&entities.Survey{},
		&entities.Question{},
		&entities.Option{},
		&entities.Answer{},
	)

	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}
