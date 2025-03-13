package handlers

import (
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/common"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func GetUser(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		//accessToken := c.Cookies("JWT")
		//if accessToken == "" {
		//	log.Println("No access token")
		//	return c.SendStatus(401)
		//}

        userId := c.Params("userId")

        // no userId provided
        if userId == "" {
        }

        token := c.Locals("jwt").(*jwt.Token)

        userId, err := common.ValidateAndReturnUserId(token, userId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        userData, err := p.GetUser(userId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(fiber.StatusInternalServerError)
        }

        jwt, err := newJWT(userId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

		return c.JSON(fiber.Map{
            "JWT": jwt,
            "user": userData,
		})
	}
}
