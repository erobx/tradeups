package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/erobx/tradeups/backend/pkg/tradeups"
)

func (p *PostgresDB) GetActiveTradeups() ([]tradeups.Tradeup, error) {
	var activeTradeups []tradeups.Tradeup
	q := `
	select t.id tradeup_id, t.rarity, t.current_status,
	coalesce(
		jsonb_agg(
			distinct jsonb_build_object(
				'username', u.username,
				'avatar', u.avatar_key
			) 
		) filter (where u.id is not null), '[]'
	) as players,
	coalesce(
		json_agg(
			json_build_object(
				'inventoryId', ts.inv_id,
				'price', (select round(cast(i.price as numeric),2)),
				'imageSrc', s.image_key
			)
		) filter (where ts.inv_id is not null), '[]'
	) as skins
	from tradeups t
	left join tradeups_skins ts on t.id = ts.tradeup_id
	left join inventory i on ts.inv_id = i.id
	left join skins s on i.skin_id = s.id
	left join users u on i.user_id = u.id
	where t.current_status = 'Active'
	group by t.id, t.rarity, t.current_status
	`
	rows, err := p.conn.Query(context.Background(), q)
	if err != nil {
		return activeTradeups, err
	}
	defer rows.Close()

    var tempTradeups []tradeups.Tradeup
    var imageKeys []string

	for rows.Next() {
		var t tradeups.Tradeup
		var playersJson, skinsJson []byte
		err := rows.Scan(&t.Id, &t.Rarity, &t.Status, &playersJson, &skinsJson)
		if err != nil {
			return activeTradeups, err
		}

		err = json.Unmarshal(playersJson, &t.Players)
		if err != nil {
			return activeTradeups, err
		}

        var tempSkins []skins.TradeupSkin
		err = json.Unmarshal(skinsJson, &tempSkins)
		if err != nil {
			return activeTradeups, err
		}

        for _, skin := range tempSkins {
            if skin.ImageSrc != "" {
                imageKeys = append(imageKeys, skin.ImageSrc)
            }
        }

        t.Skins = tempSkins

        tempTradeups = append(tempTradeups, t)
	}

    if err := rows.Err(); err != nil {
        return activeTradeups, err
    }

    urlMap := p.urlManager.GetUrls(imageKeys)

    for i := range tempTradeups {
        for j := range tempTradeups[i].Skins {
            if url, exists := urlMap[tempTradeups[i].Skins[j].ImageSrc]; exists {
                tempTradeups[i].Skins[j].ImageSrc = url
            }
        }
        activeTradeups = append(activeTradeups, tempTradeups[i])
    }

	return activeTradeups, nil 
}

func (p *PostgresDB) GetTradeup(id string) (tradeups.Tradeup, error) {
    var t tradeups.Tradeup
    q := `
    select t.id tradeup_id, t.rarity, t.current_status,
	coalesce(
		jsonb_agg(
			distinct jsonb_build_object(
				'username', u.username,
				'avatar', u.avatar_key
			) 
		) filter (where u.id is not null), '[]'
	) as players,
	coalesce(
		json_agg(
			json_build_object(
				'inventoryId', ts.inv_id,
                'userId', u.id,
				'price', (select round(cast(i.price as numeric),2)),
                'name', s.name,
                'wear', i.wear_str,
                'skinFloat', i.wear_num,
                'isStatTrak', i.is_stattrak,
				'imageSrc', s.image_key,
                'createdAt', to_char(i.created_at, 'YYYY/MM/DD HH12:MI:SS')
			)
		) filter (where ts.inv_id is not null), '[]'
	) as skins
	from tradeups t
	left join tradeups_skins ts on t.id = ts.tradeup_id
	left join inventory i on ts.inv_id = i.id
	left join skins s on i.skin_id = s.id
	left join users u on i.user_id = u.id
	where t.id = $1
    group by t.id
    `

    var playersJson, skinsJson []byte
    var imageKeys []string
    row := p.conn.QueryRow(context.Background(), q, id)
    err := row.Scan(&t.Id, &t.Rarity, &t.Status, &playersJson, &skinsJson)
    if err != nil {
        return t, err
    }

    err = json.Unmarshal(playersJson, &t.Players)
    if err != nil {
        return t, err
    }

    err = json.Unmarshal(skinsJson, &t.Skins)
    if err != nil {
        return t, err
    }

    for _, skin := range t.Skins {
        if skin.ImageSrc != "" {
            imageKeys = append(imageKeys, skin.ImageSrc)
        }
    }

    urlMap := p.urlManager.GetUrls(imageKeys)

    for i := range t.Skins {
        if url, exists := urlMap[t.Skins[i].ImageSrc]; exists {
            t.Skins[i].ImageSrc = url
        }
    }

    return t, nil
}

func (p *PostgresDB) TradeupIsFull(tradeupId string) error {
    var numSkins int
    q := "select count(tradeup_id) from tradeups_skins where tradeup_id=$1"
    row := p.conn.QueryRow(context.Background(), q, tradeupId)
    err := row.Scan(&numSkins)
    if err != nil {
        return err
    }

    if numSkins > 10 {
        return fmt.Errorf("Tradeup full")
    }
    return nil
}

func (p *PostgresDB) AddSkinToTradeup(userId, tradeupId string, invId int) error {
    // if the user actually owns the skin to add
    var exists bool
    q := "select exists(select 1 from inventory where user_id=$1 and id=$2)"
    row := p.conn.QueryRow(context.Background(), q, userId, invId)
    err := row.Scan(&exists)
    if err != nil {
        return err
    }
    
    if !exists {
        return fmt.Errorf("User does not own that item")
    }

    // finally add the skin
    q = "insert into tradeups_skins values($1,$2)"
    tag, err := p.conn.Exec(context.Background(), q, tradeupId, invId)
    if err != nil {
        return err
    }

    if !tag.Insert() {
        return fmt.Errorf("Not an insert statement")
    }

    return nil
}

func (p *PostgresDB) RemoveSkinFromTradeup(tradeupId string, invId int) (skins.InventorySkin, error) {
    var invSkin skins.InventorySkin
    var imageKey string
    q := `
    with deleted_skin as (
        delete from tradeups_skins ts
        where tradeup_id=$1 and inv_id=$2
        returning inv_id
    )
    select i.id, i.wear_str, i.wear_num, round(cast(i.price as numeric),2), i.is_stattrak, to_char(i.created_at, 'YYYY/MM/DD HH12:MI:SS'),
		s.name, s.rarity, s.collection, s.image_key
    from inventory i
	join skins s on s.id = i.skin_id
	where i.id=$3
    order by s.image_key, i.wear_str
    `
    row := p.conn.QueryRow(context.Background(), q, tradeupId, invId, invId)
    err := row.Scan(&invSkin.Id, &invSkin.Wear, &invSkin.SkinFloat, &invSkin.Price, &invSkin.IsStatTrak, &invSkin.CreatedAt,
        &invSkin.Name, &invSkin.Rarity, &invSkin.Collection, &imageKey)
    
    if err != nil {
        return invSkin, err
    }

    urlMap := p.urlManager.GetUrls([]string{imageKey})
    invSkin.ImageSrc = urlMap[imageKey]
    return invSkin, err
}
