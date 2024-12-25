package main

import (
	"log"
	"net/http"
	"os"

	"github.com/agrotention/backend/handlers"
	"github.com/agrotention/backend/services"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// DB Connection
	dbUri := os.Getenv("DB_URI")
	db, err := gorm.Open(mysql.Open(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	// Create Mux
	mux := http.NewServeMux()

	// Init Services and Handlers
	// == Users
	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)
	userHandler.RegisterRouter(mux)

	// Run Server
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
