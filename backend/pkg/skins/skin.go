package skins

type Skin struct {
	Name       string  `json:"name"`
	WeaponName string  `json:"weaponName"`
	Wear       string  `json:"wear"`
	Color      string  `json:"color"`
	Collection string  `json:"collection"`
	MinFloat   float64 `json:"minFloat"`
	MaxFloat   float64 `json:"maxFloat"`
}

func NewSkin(name, weaponName, wear, color, collection string, minFloat, maxFloat float64) Skin {
	return Skin{
		Name:       name,
		WeaponName: weaponName,
		Wear:       wear,
		Color:      color,
		Collection: collection,
		MinFloat:   minFloat,
		MaxFloat:   maxFloat,
	}
}

// Should read all available skins in CS2 from steam api or file
// and insert them into database.
func Insert() {

}
