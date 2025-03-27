package user

import (
	"time"

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

type UserData struct {
    Id string `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
    AvatarSrc string `json:"avatarSrc"`
    RefreshTokenVersion int `json:"refreshTokenVersion"`
    CreatedAt time.Time `json:"createdAt"`
    Balance float64 `json:"balance"`
}

type Inventory struct {
	Skins []skins.InventorySkin `json:"skins"`
}

type Stats struct {
    RecentWinnings []skins.InventorySkin `json:"recentWinnings"`
    TradeupsEntered int `json:"tradeupsEntered"`
    TradeupsWon int `json:"tradeupsWon"`
}

type RegisteredUserPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type EmailPayload struct {
	Email string `json:"email"`
}
