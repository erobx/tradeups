package db

import (
	"context"
	"slices"

	"github.com/erobx/tradeups/backend/pkg/skins"
)

func (p *PostgresDB) BuyCrate(userId, name string, count int) ([]skins.InventorySkin, error) {
    var newSkins []skins.InventorySkin
    var imageKeys []string

    q := `
    SELECT oc.id, oc.skin_id, oc.wear_str, oc.wear_num, round(cast(oc.price as numeric),2), oc.is_stattrak, to_char(oc.created_at, 'YYYY/MM/DD HH12:MI:SS'),
            s.name, s.rarity, s.collection, s.image_key
    FROM open_crate(
        (SELECT id FROM users WHERE id=$1),
        $2,
        $3
    ) as oc
    join skins s on oc.skin_id = s.id
    `
    rows, err := p.conn.Query(context.Background(), q, userId, name, count)
    if err != nil {
        return newSkins, err
    }

    tempItems := make(map[string][]skins.InventorySkin)
	for rows.Next() {
		var s skins.InventorySkin
        var skinId int
		var imageKey string

		err := rows.Scan(&s.Id, &skinId, &s.Wear, &s.SkinFloat, &s.Price,
                        &s.IsStatTrak, &s.CreatedAt, &s.Name, &s.Rarity, &s.Collection,
                        &imageKey)
		if err != nil {
			return newSkins, err
		}

        if !slices.Contains(imageKeys, imageKey) {
            imageKeys = append(imageKeys, imageKey)
        }

        tempItems[imageKey] = append(tempItems[imageKey], s)
	}

    urlMap := p.urlManager.GetUrls(imageKeys)

    for imageKey, skinGroup := range tempItems {
        url, exists := urlMap[imageKey]
        if !exists {
            continue
        }

        for i := range skinGroup {
            skinGroup[i].ImageSrc = url
        }

        newSkins = append(newSkins, skinGroup...)
    }

	return newSkins, rows.Err()
}
