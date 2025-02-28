package handlers

import (
	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/gofiber/fiber/v3"
)

func GetActiveTradeups(db *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {

        return nil
    }
}
