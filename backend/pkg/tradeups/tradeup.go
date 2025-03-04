package tradeups

import "github.com/erobx/tradeups/backend/pkg/skins"

type Tradeup struct {
    Id int `json:"id"`
    Rarity string `json:"rarity"`
    Status string `json:"status"`
    Skins []skins.TradeupSkin `json:"skins"`
    Players []Player `json:"players"`
}

type Player struct {
    Username string `json:"username"`
    Avatar string `json:"avatar"`
}
