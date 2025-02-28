package handlers

import (
	"fmt"
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/pkg/common"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// {id: 0, name: "M4A4 | Howl", wear: "Factory New", rarity: "Contraband", float: 0.01, isStatTrak: true, imgSrc: "/m4a4-howl.png"},

func GetInventory(p *db.PostgresDB) fiber.Handler {
	return func(c fiber.Ctx) error {
		urlUserId := c.Params("id")
		// for now using Bearer token instead of jwt in the cookie bc of localhost

		reqHeaders := c.GetReqHeaders()
		authHeader, ok := reqHeaders["Authorization"]
		if !ok {
			return c.SendStatus(403)
		}

		// Bearer jwt
		tokenString := authHeader[0][7:]

		// verify jwt
		token, err := verifyJwt(tokenString)
		if err != nil {
			log.Println(err)
			c.SendStatus(403)
		}

		var jwtUserId string
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			jwtUserId, _ = claims.GetSubject()
		} else {
			fmt.Println(err)
			return c.SendStatus(500)
		}

		if jwtUserId != urlUserId {
			return c.SendStatus(403)
		}

		inv, err := p.GetInventory(jwtUserId)
		if err != nil {
            log.Println(err)
			return c.SendStatus(500)
		}

		return c.JSON(inv)
	}
}

func verifyJwt(tokenString string) (*jwt.Token, error) {
	keyBytes, err := common.ReadPubKey()
	if err != nil {
		return nil, err
	}

	verifyingKey, err := jwt.ParseECPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return verifyingKey, err
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
