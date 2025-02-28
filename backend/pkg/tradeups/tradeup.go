package tradeups

import "github.com/erobx/tradeups/backend/pkg/skins"

type Tradeup struct {
    Id int
    Rarity string
    Status string
    Skins []skins.TradeupSkin
    Players []Player
}

type Player struct {
    Username string
    Avatar string
}
