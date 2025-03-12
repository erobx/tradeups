package user

import (
	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/google/uuid"
)

type User struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Hash     string    `json:"hash"`
}

type Inventory struct {
	Skins []skins.InventorySkin `json:"skins"`
}

type RegisteredUserPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type EmailPayload struct {
	Email string `json:"email"`
}
