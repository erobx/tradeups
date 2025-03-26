package handlers

import (
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
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
        token := c.Locals("jwt").(*jwt.Token)

		var jwtUserId string
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			jwtUserId, _ = claims.GetSubject()
		} else {
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
        
        //if count, _ := userCache[jwtUserId]; count > 5 {
        //    log.Println("user has added more than 5 skins")
        //    return nil
        //}
        // check if item user is trying to add belongs to them
        err := p.TradeupIsFull(payload.TradeupId)
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

func RemoveSkinFromTradeup(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        token := c.Locals("jwt").(*jwt.Token)

        var userId string
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userId, _ = claims.GetSubject()
		} else {
			return c.SendStatus(500)
		}        

        payload := new(TradeupSkinsPayload)
        if err := c.Bind().Body(payload); err != nil {
            log.Println(err)
            return err
        }

        // check if item user is trying to remove belongs to them
        exists := p.IsUsersSkin(userId, payload.InvId)
        if !exists {
            return c.SendStatus(fiber.StatusForbidden)
        }

        returnedSkin, err := p.RemoveSkinFromTradeup(payload.TradeupId, payload.InvId)
        if err != nil {
            log.Println(err)
            return c.SendStatus(500)
        }

        return c.JSON(returnedSkin)
    }
}

type TradeupPayload struct {
    Rarity string `json:"rarity"`
}
// admin handler
func NewTradeup(p *db.PostgresDB) fiber.Handler {
    return func(c fiber.Ctx) error {
        payload := new(TradeupPayload)
        c.Bind().Body(payload)
        log.Printf("Creating new tradeup of type %s\n", payload.Rarity)

        err := p.CreateTradeup(payload.Rarity)
        if err != nil {
            return c.SendString("Could not create new tradeup")
        }

        return c.SendString("Created new tradeup")
    }
}
