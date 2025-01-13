package main

import (
	"log"
	"os"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/db"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db.Init()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	router.SetupUserRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
