package handlers

import (
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
    "github.com/erobx/tradeups/backend/pkg/common"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type TradeupSkinsPayload struct {
    InvId int `json:"invId"`
    TradeupId string `json:"tradeupId"`
}


var userCache = make(map[string]int)

func AddSkinToTradeup(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
		// for now using Bearer token instead of jwt in the cookie bc of localhost
		token, err := common.ValidateHeaders(c)
        if err != nil {
            return err
        }

		var jwtUserId string
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			jwtUserId, _ = claims.GetSubject()
		} else {
			log.Println(err)
			return c.SendStatus(500)
		}

        if _, ok := userCache[jwtUserId]; ok {
            userCache[jwtUserId] += 1
        } else {
            userCache[jwtUserId] = 1
        }

        payload := new(TradeupSkinsPayload)
        if err := c.Bind().Body(payload); err != nil {
            log.Println(err)
            return err
        }
        
        if count, _ := userCache[jwtUserId]; count > 5 {
            log.Println("user has added more than 5 skins")
            return nil
        }
        // check if item user is trying to add belongs to them
        err = p.TradeupIsFull(payload.TradeupId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

        err = p.AddSkinToTradeup(jwtUserId, payload.TradeupId, payload.InvId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }
		
        return c.SendStatus(201)
    }
}
