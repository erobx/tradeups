package handlers

import (
	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/gofiber/fiber/v3"
)

// { exists: bool }
func CheckEmail(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"exists": true,
		})
	}
}
