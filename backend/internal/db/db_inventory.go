package db

import (
	"context"
	"fmt"
	"slices"

	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/erobx/tradeups/backend/pkg/user"
)

func (p *PostgresDB) GetInventory(userId string) (user.Inventory, error) {
	var inv user.Inventory
	var items []skins.InventorySkin
    var imageKeys []string

	q :=`
    WITH unique_image_skins AS (
        SELECT DISTINCT s.image_key
        FROM inventory i 
        JOIN skins s ON s.id = i.skin_id
        WHERE i.user_id = $1
        AND NOT EXISTS (
            SELECT 1 FROM tradeups_skins ts
            WHERE ts.inv_id = i.id
        )
    )
    select i.id, i.wear_str, i.wear_num, round(cast(i.price as numeric),2), i.is_stattrak, i.created_at,
		s.name, s.rarity, s.collection, s.image_key
	from inventory i
	join skins s on s.id = i.skin_id
	where i.user_id=$1
		and not exists (
			select 1 from tradeups_skins ts
			where ts.inv_id = i.id
		)
    order by s.image_key, i.wear_str
	`

	rows, err := p.conn.Query(context.Background(), q, userId)
	if err != nil {
		return inv, err
	}
	defer rows.Close()

    tempItems := make(map[string][]skins.InventorySkin)
	for rows.Next() {
		var s skins.InventorySkin
		var imageKey string

		err := rows.Scan(&s.Id, &s.Wear, &s.SkinFloat, &s.Price,
                        &s.IsStatTrak, &s.CreatedAt, &s.Name, &s.Rarity, &s.Collection,
                        &imageKey)
		if err != nil {
			return inv, err
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

        items = append(items, skinGroup...)
    }

    if len(items) == 0 {
        return inv, fmt.Errorf("items empty")
    }

	inv.Skins = items
	return inv, rows.Err()
}

func (p *PostgresDB) DeleteSkin(userId, skinId string) error {
    q := "delete from inventory where user_id=$1 and id=$2"
    tag, err := p.conn.Exec(context.Background(), q, userId, skinId)
    if err != nil {
        return err
    }

    if !tag.Delete() {
        return fmt.Errorf("Not a delete statement")
    }

    return nil
}
