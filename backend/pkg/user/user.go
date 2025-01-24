package user

import (
	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/google/uuid"
)

type User struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Hash     string    `json:"hash"`
}

type Inventory struct {
	Skins []skins.Skin
}
