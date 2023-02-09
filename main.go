package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/oeggy03/h4g-backend/connect"
	"github.com/oeggy03/h4g-backend/routes"
)

func main() {
	connect.Connect()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env files - port")
	} else {
		log.Println("Connected successfully")
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	//to prevent Cors issues
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	//set up routes
	routes.Setup(app)

	//port retrieved from env. currently running on port 3001 as frontend is on 3000
	app.Listen(":" + port)

}
