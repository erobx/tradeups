package tradeups

import (
	"time"

	"github.com/erobx/tradeups/backend/pkg/skins"
)

type Tradeup struct {
    Id int `json:"id"`
    Rarity string `json:"rarity"`
    Status string `json:"status"`
    StopTime time.Time `json:"stopTime"`
    Skins []skins.TradeupSkin `json:"skins"`
    Players []Player `json:"players"`
}

type Player struct {
    Username string `json:"username"`
    Avatar string `json:"avatar"`
}
