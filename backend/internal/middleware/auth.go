package middleware

import (
    "fmt"
	"log"

	"github.com/erobx/tradeups/backend/pkg/common"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
    return func(c fiber.Ctx) error {
        reqHeaders := c.GetReqHeaders()
        authHeader, ok := reqHeaders[fiber.HeaderAuthorization]
        if !ok {
            return c.SendStatus(fiber.StatusBadRequest)
        }

        tokenString := authHeader[0][7:]
        token, err := verifyJwt(tokenString)
        if err != nil {
            log.Println(err)
            return c.SendStatus(fiber.StatusUnauthorized)
        }
        c.Locals("jwt", token)
        return c.Next()
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

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
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
