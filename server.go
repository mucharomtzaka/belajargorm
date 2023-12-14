package main

import (
	"belajargo/Config"
	"belajargo/Models"
	"belajargo/Routes"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	//load env file
	env := godotenv.Load()

	if env != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	Config.DatabaseInit()
	defer Config.GetDB().DB()

	// Perform migrations using AutoMigrate
	db := Config.GetDB()

	err := db.AutoMigrate(&Models.Course{})
	if err != nil {
		panic(err)
	}

	// Set up Routes
	Routes.SetupRoutes(e)

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
