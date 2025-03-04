package handlers

import (
    "log"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/gofiber/fiber/v3"
)

func GetActiveTradeups(db *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        res, err := db.GetActiveTradeups()
        if err != nil {
            log.Println(err)
            c.SendStatus(500)
        }

        return c.JSON(res)
    }
}
