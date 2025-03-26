package db

import (
	"context"
	"fmt"
	"slices"

	"github.com/erobx/tradeups/backend/pkg/skins"
)

func (p *PostgresDB) BuyCrate(userId, name, rarity string, count int) ([]skins.InventorySkin, float64, error) {
    var newSkins []skins.InventorySkin
    var newBalance float64
    var imageKeys []string

    // check if user has enough funds to buy the crate
    q := `
    with crate_details as (
        select cost from crates where name = $1
    ),
    balance_check as (
        select balance >= (select cost from crate_details) as sufficient_funds
        from users where id = $2
    )
    update users set balance = balance - (select cost from crate_details)
    where id = $2 and exists(
        select 1 from balance_check where sufficient_funds = true
    )
    `
    tag, err := p.conn.Exec(context.Background(), q, name, userId)
    if err != nil {
        return newSkins, newBalance, err
    }
    // no rows were affected => insufficient funds
    if tag.RowsAffected() == 0 {
        return newSkins, newBalance, fmt.Errorf("Couldn't buy crate: insufficient funds")
    }

    // otherwise get new balance
    q = "select balance from users where id = $1"
    err = p.conn.QueryRow(context.Background(), q, userId).Scan(&newBalance)
    if err != nil {
        return newSkins, newBalance, err
    }

    // TODO: change when I add certain skins to a crate (crate_skins)
    // right now still adds XX random skins of the selected rarity
    // now open the selected crate
    q = `
    SELECT oc.id, oc.skin_id, oc.wear_str, oc.wear_num, round(cast(oc.price as numeric),2), oc.is_stattrak, oc.created_at,
            s.name, s.rarity, s.collection, s.image_key
    FROM open_crate(
        (SELECT id FROM users WHERE id=$1),
        $2,
        $3
    ) as oc
    join skins s on oc.skin_id = s.id
    `
    rows, err := p.conn.Query(context.Background(), q, userId, rarity, count)
    if err != nil {
        return newSkins, newBalance, err
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
			return newSkins, newBalance, err
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

	return newSkins, newBalance, rows.Err()
}
