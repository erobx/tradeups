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

func GetNextRarity(prevRarity string) string {
    if prevRarity == "Contraband" {
        return ""
    }

    switch prevRarity {
    case "Consumer":
        return "Industrial"
    case "Industrial":
        return "Mil-Spec"
    case "Mil-Spec":
        return "Restricted"
    case "Restricted":
        return "Classified"
    case "Classified":
        return "Covert"
    default:
        return "Contraband"
    }
}

func GetWearNameFromFloat(wear float64) string {
    if wear < 0.07 {
        return "Factory New"
    } else if wear < 0.15 {
        return "Minimal Wear"
    } else if wear < 0.38 {
        return "Field-Tested"
    } else if wear < 0.45 {
        return "Well-Worn"
    } else {
        return "Battle-Scarred"
    }
}
