package api

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/yuri7030/final-project/internal/api/config"
	"github.com/yuri7030/final-project/internal/api/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dbName = config.GetDbName()
	fmt.Println("dbName", dbName)
	var dbPath string = fmt.Sprintf("data/%s.db", dbName)
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("database file not exist")
		f, err := os.Create(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

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
