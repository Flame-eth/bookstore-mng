package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func SetupDatabase() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Read the DSN from the environment variable
	dsn := os.Getenv("DSN")
	if dsn == "" {
		fmt.Println("DSN environment variable not set")
		return
	}
	// Open the connection to the database
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
