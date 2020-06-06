// models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=belajargo password=root sslmode=disable")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Twitter{})

	DB = database

}
