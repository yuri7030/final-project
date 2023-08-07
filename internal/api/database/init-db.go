package database

import (
	"fmt"

	"github.com/yuri7030/final-project/internal/api/config"
	"github.com/yuri7030/final-project/internal/api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dbName = config.GetValue("DB_NAME")
	var dbHost = config.GetValue("DB_HOST")
	var dbUser = config.GetValue("DB_USER")
	var dBPassword = config.GetValue("DB_PASSWORD")
	var dbPort = config.GetValue("DB_PORT")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dBPassword, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{
		// LogLevel: 2
	})

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
