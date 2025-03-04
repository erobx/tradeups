package skins

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
}

type TradeupSkin struct {
	InventoryId int `json:"inventoryId"`
	Price float64 `json:"price"`
	ImageSrc string `json:"imageSrc"`
}

func NewSkin(name, rarity, collection string, minFloat, maxFloat float64) Skin {
	return Skin{
		Name:       name,
		Rarity:      rarity,
		Collection: collection,
		MinFloat:   minFloat,
		MaxFloat:   maxFloat,
	}
}

func NewInventorySkin(id int, sf, p float64, isSt bool, n, w, r, col, imgSrc string) InventorySkin {
	return InventorySkin{
		Id: id,
		SkinFloat: sf,
		Price: p,
		IsStatTrak: isSt,
		Name: n,
		Wear: w,
		Rarity: r,
		Collection: col,
		ImageSrc: imgSrc,
	}
}
