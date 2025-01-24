package main

import (
	"context"
	"fmt"
	"log"

	"github.com/erobx/tradeups/backend/pkg/db"
	"github.com/erobx/tradeups/backend/pkg/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func defineRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/getUser", handlers.GetUser)
	api.Post("/createUser", handlers.CreateUser)
}

func main() {
	fmt.Println("Starting server...")

	app := fiber.New()
	db.Connect()
	defer db.Postgresql.Close(context.Background())

	//app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type"},
	}))

	defineRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
