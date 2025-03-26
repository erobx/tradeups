package skins

import "time"

type Skin struct {
	Name       string  `json:"name"`
	Rarity      string  `json:"rarity"`
	Collection string  `json:"collection"`
	MinFloat   float64 `json:"minFloat"`
	MaxFloat   float64 `json:"maxFloat"`
	CanBeStatTrak bool `json:"canBeStatTrak"`
	ImageKey string `json:"imageKey"`
}

type InventorySkin struct {
	Id int `json:"id"`
	SkinFloat float64 `json:"skinFloat"`
	Price float64 `json:"price"`
	IsStatTrak bool `json:"isStatTrak"`
	Name string `json:"name"`
	Wear string `json:"wear"`
	Rarity string `json:"rarity"`
	Collection string `json:"collection"`
	ImageSrc string `json:"imageSrc"`
    CreatedAt time.Time `json:"createdAt"`
}

type TradeupSkin struct {
	InventoryId int `json:"inventoryId"`
    UserId string `json:"userId"`
	Price float64 `json:"price"`
    Name string `json:"name"`
    Wear string `json:"wear"`
    SkinFloat float64 `json:"skinFloat"`
    IsStatTrak bool `json:"isStatTrak"`
	ImageSrc string `json:"imageSrc"`
    CreatedAt string `json:"createdAt"`
}
