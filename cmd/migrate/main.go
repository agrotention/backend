package main

import (
	"log"
	"os"

	"github.com/agrotention/backend/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbUri := os.Getenv("DB_URI")
	db, err := gorm.Open(mysql.Open(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{}, &models.UserInfo{}, &models.UserRole{})

}
