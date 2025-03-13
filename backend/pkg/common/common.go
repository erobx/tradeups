package common

import (
    "fmt"
	"os"
	"strings"

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
