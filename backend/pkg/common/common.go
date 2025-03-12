package common

import (
    "fmt"
    "log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func ReadPrivKey() ([]byte, error) {
	b, err := os.ReadFile(os.Getenv("PRIVATE_KEY"))
	return b, err
}

func ReadPubKey() ([]byte, error) {
	b, err := os.ReadFile(os.Getenv("PUBLIC_KEY"))
	return b, err
}

func PrefixKey(key string) string {
    before, _, found := strings.Cut(key, "-")
    if !found {
        return key
    }

    key = "guns/" + before + "/" + key
    return key
}

func ValidateHeaders(c fiber.Ctx) (*jwt.Token, error) {
    reqHeaders := c.GetReqHeaders()
    authHeader, ok := reqHeaders["Authorization"]
    if !ok {
        return nil, c.SendStatus(403)
    }

    tokenString := authHeader[0][7:]

    token, err := verifyJwt(tokenString)
    if err != nil {
        log.Println(err)
        c.SendStatus(403)
    }

    return token, nil
}

func ValidateAndReturnUserId(token *jwt.Token, userId string) (jwtUserId string, err error) {
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        jwtUserId, _ = claims.GetSubject()
    } else {
        err = fmt.Errorf("couldn't validate jwt")
        return jwtUserId, err
    }

    if jwtUserId != userId {
        err = fmt.Errorf("unauthorized user")
        return jwtUserId, err
    }
    return
}

func verifyJwt(tokenString string) (*jwt.Token, error) {
	keyBytes, err := ReadPubKey()
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
