package main

import (
	"blog-be-go/database"
	"blog-be-go/routes"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	database.Connect()
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
