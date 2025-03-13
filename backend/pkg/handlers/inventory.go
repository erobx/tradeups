package handlers

import (
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/common"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// {id: 0, name: "M4A4 | Howl", wear: "Factory New", rarity: "Contraband", float: 0.01, isStatTrak: true, imgSrc: "/m4a4-howl.png"},
func GetInventory(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		urlUserId := c.Params("userId")
		// for now using Bearer token instead of jwt in the cookie bc of localhost
        token := c.Locals("jwt").(*jwt.Token)

        jwtUserId, err := common.ValidateAndReturnUserId(token, urlUserId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

		inv, err := p.GetInventory(jwtUserId)
		if err != nil {
            log.Println(err)
			return c.JSON(fiber.Map{
                "skins": "empty",
            })
		}

		return c.JSON(inv)
	}
}

func DeleteSkin(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        urlUserId := c.Params("userId")
        urlInvId := c.Params("invId")

        token := c.Locals("jwt").(*jwt.Token)

        jwtUserId, err := common.ValidateAndReturnUserId(token, urlUserId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

        err = p.DeleteSkin(jwtUserId, urlInvId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

        log.Printf("User: %s deleted item %s from their inventory\n", urlUserId, urlInvId)
        return c.SendStatus(204)
    }
}


