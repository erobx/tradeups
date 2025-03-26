package handlers

import (
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type CratePayload struct {
    Name string `json:"name"`
    Rarity string `json:"rarity"`
    Count int `json:"count"`
}

type SuccessfulPurchase struct {
    Skins []skins.InventorySkin `json:"skins"`
    Balance float64 `json:"balance"`
}

func BuyCrate(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        token := c.Locals("jwt").(*jwt.Token)

		var jwtUserId string
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			jwtUserId, _ = claims.GetSubject()
		} else {
			return c.SendStatus(500)
		}

        log.Println("Buying crate...")
        payload := new(CratePayload)
        if err := c.Bind().Body(payload); err != nil {
            log.Println(err)
            return err
        }

        newSkins, newBalance, err := p.BuyCrate(jwtUserId, payload.Name, payload.Rarity, payload.Count)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

        return c.JSON(SuccessfulPurchase{
            newSkins,
            newBalance,
        })
    }
}
