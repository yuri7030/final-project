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
	var dbName = config.GetDbName()
	var dbUser = config.GetDbUser()
	var dBPassword = config.GetDbPassword()
	var dbPort = config.GetDbPort()

	dbPath := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dBPassword, dbName, dbPort)
	fmt.Println("dbPath", dbPath)

	database, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})

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
		return
	}

	DB = database
}
