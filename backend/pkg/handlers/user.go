package handlers

import (
	"log"
	"strings"

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
        token := c.Locals("jwt").(*jwt.Token)

        // no userId provided
        if userId == "" {
            if claims, ok := token.Claims.(jwt.MapClaims); ok {
                userId, _ = claims.GetSubject()
            } else {
                return c.SendStatus(fiber.StatusUnauthorized)
            }
            return sendUserData(p, c, userId)
        } else {
            userId, err := common.ValidateAndReturnUserId(token, userId)
            if err != nil {
                log.Println(err)
                return c.SendStatus(fiber.StatusUnauthorized)
            }

            return sendUserData(p, c, userId)
        }
	}
}

func GetUserStats(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        userId := c.Params("userId")
        token := c.Locals("jwt").(*jwt.Token)

        userId, err := common.ValidateAndReturnUserId(token, userId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        stats, err := p.GetStats(userId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(fiber.StatusInternalServerError)
        }

        return c.JSON(stats)
    }
}

func GetRecentTradeups(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        userId := strings.TrimSpace(c.Params("userId"))
        token := c.Locals("jwt").(*jwt.Token)

        userId, err := common.ValidateAndReturnUserId(token, userId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        recentTradeups, err := p.GetRecentTradeups(userId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(fiber.StatusInternalServerError)
        }

        return c.JSON(recentTradeups)
    }
}

func sendUserData(p *db.PostgresDB, c fiber.Ctx, userId string) error {
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
