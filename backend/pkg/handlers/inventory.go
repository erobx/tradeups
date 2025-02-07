package handlers

import (
	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/gofiber/fiber/v3"
)

func GetSkins(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"test": "works",
		})
	}
}
