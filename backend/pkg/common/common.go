package common

import (
	"os"
    "strings"
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
