package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Starting server...")
	app := fiber.New()

	app.Get("/api/test", func(c fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.JSON(fiber.Map{
			"test": "works",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
