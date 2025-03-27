package db

import (
	"context"
	"encoding/json"
	"log"
	"slices"
	"time"

	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/erobx/tradeups/backend/pkg/tradeups"
	"github.com/erobx/tradeups/backend/pkg/user"
)

func (p *PostgresDB) FindEmail(email string) (bool, error) {
	var exists bool
	q := "select exists(select 1 from users where email=$1 limit 1)"
	row := p.conn.QueryRow(context.Background(), q, email)
	row.Scan(&exists)

	return exists, nil
}

func (p *PostgresDB) FindUsername(username string) (bool, error) {
	var exists bool
	q := "select exists(select 1 from users where username=$1 limit 1)"
	row := p.conn.QueryRow(context.Background(), q, username)
	row.Scan(&exists)

	return exists, nil
}

func (p *PostgresDB) CreateUser(u *user.User) (user.UserData, error) {
    var userData user.UserData
    var avatarKey string
	q := `
    insert into users(id, username, email, hash, created_at) values($1,$2,$3,$4,$5)
    returning id, username, email, avatar_key, refresh_token_version,
    created_at, balance
    `
	row := p.conn.QueryRow(context.Background(), q, u.Uuid, u.Username, u.Email, u.Hash, time.Now())
    err := row.Scan(&userData.Id, &userData.Username, &userData.Email, &avatarKey, &userData.RefreshTokenVersion, &userData.CreatedAt, &userData.Balance)

    if avatarKey != "" {
        urlMap := p.urlManager.GetUrls([]string{avatarKey})
        if key, ok := urlMap[avatarKey]; ok {
            avatarKey = key
        }
    }
    userData.AvatarSrc = avatarKey

    // give user random skins
    _, _, err = p.BuyCrate(u.Uuid.String(), "ConOG", "Consumer", 8)
    if err != nil {
        log.Printf("Could not give user %s new skins\n", u.Uuid.String())
    }

	return userData, err
}

func (p *PostgresDB) GetHash(email string) (id, hash string, err error) {
	q := "select id, hash from users where email=$1"
	row := p.conn.QueryRow(context.Background(), q, email)
	err = row.Scan(&id, &hash)
	if err != nil {
		return id, hash, err
	}

	return id, hash, err
}

func (p *PostgresDB) GetUser(id string) (user.UserData, error) {
    var userData user.UserData
    var avatarKey string
    q := `
    select id, username, email, avatar_key, refresh_token_version,
    created_at, balance from users where
    id = $1
    `
    row := p.conn.QueryRow(context.Background(), q, id)
    err := row.Scan(&userData.Id, &userData.Username, &userData.Email, &avatarKey, &userData.RefreshTokenVersion, &userData.CreatedAt, &userData.Balance)

    if avatarKey != "" {
        urlMap := p.urlManager.GetUrls([]string{avatarKey})
        if key, ok := urlMap[avatarKey]; ok {
            avatarKey = key
        }
    }
    userData.AvatarSrc = avatarKey

	return userData, err
}

func (p *PostgresDB) IsUsersSkin(userId string, invId int) bool {
    var exists bool
    q := "select exists(select 1 from inventory where user_id=$1 and id=$2)"
    row := p.conn.QueryRow(context.Background(), q, userId, invId)
    row.Scan(&exists)
    return exists
}

func (p *PostgresDB) GetStats(userId string) (user.Stats, error) {
    var stats user.Stats
    var winnings []skins.InventorySkin
    var imageKeys []string

    // recent winnings, tradeups entered, tradeups won
    q := `
    select i.id, i.wear_str, i.wear_num, round(cast(i.price as numeric),2), i.is_stattrak, i.created_at,
		s.name, s.rarity, s.collection, s.image_key
    from inventory i
    join skins s on s.id = i.skin_id
    where i.user_id=$1 and i.was_won=true
    order by i.created_at desc limit 4 
    `
    rows, err := p.conn.Query(context.Background(), q, userId)
    if err != nil {
        return stats, err
    }
    defer rows.Close()

    for rows.Next() {
        var s skins.InventorySkin
        var imageKey string
        err := rows.Scan(&s.Id, &s.Wear, &s.SkinFloat, &s.Price, &s.IsStatTrak,
                        &s.CreatedAt, &s.Name, &s.Rarity, &s.Collection, &imageKey)
        if err != nil {
            return stats, err
        }

        if !slices.Contains(imageKeys, imageKey) {
            imageKeys = append(imageKeys, imageKey)
        }
        winnings = append(winnings, s)
    }

    if len(imageKeys) > 0 {
        urlMap := p.urlManager.GetUrls(imageKeys)
        for i := range winnings {
            url, exists := urlMap[imageKeys[i]]
            if exists {
                winnings[i].ImageSrc = url
            }
        }
    }

    q = `
    select count(distinct t.id) as tradeups_entered from tradeups t
    join tradeups_skins ts on ts.tradeup_id = t.id
    join inventory i on i.id = ts.inv_id
    where i.user_id = $1 and t.current_status = 'Completed'
    `
    err = p.conn.QueryRow(context.Background(), q, userId).Scan(&stats.TradeupsEntered)
    if err != nil {
        return stats, err
    }

    q = `
    select count(*) as tradeups_won from tradeups
    where winner = $1 and current_status = 'Completed'
    `
    err = p.conn.QueryRow(context.Background(), q, userId).Scan(&stats.TradeupsWon)
    if err != nil {
        return stats, err
    }

    stats.RecentWinnings = winnings
    return stats, rows.Err()
}

func (p *PostgresDB) GetRecentTradeups(userId string) ([]tradeups.Tradeup, error) {
    var recentTradeups []tradeups.Tradeup

    q := `
    WITH TradeupDetails AS (
    SELECT 
        t.id AS tradeup_id,
        t.rarity,
        t.current_status AS status,
        COALESCE(SUM(i.price), 0) AS total_value
    FROM 
        tradeups t
    JOIN 
        tradeups_skins ts ON t.id = ts.tradeup_id
    JOIN 
        inventory i ON ts.inv_id = i.id
    WHERE
        i.user_id = $1
    GROUP BY 
        t.id, t.rarity, t.current_status
    )
    SELECT 
        td.tradeup_id AS id,
        td.rarity,
        td.status,
        td.total_value AS value,
        json_agg(
            json_build_object(
                'inventoryId', i.id,
                'userId', i.user_id,
                'price', i.price,
                'name', s.name,
                'wear', i.wear_str,
                'skinFloat', i.wear_num,
                'isStatTrak', i.is_stattrak,
                'imageSrc', s.image_key,
                'createdAt', i.created_at::text
            )
        ) AS skins
    FROM 
        TradeupDetails td
    JOIN 
        tradeups_skins ts ON td.tradeup_id = ts.tradeup_id
    JOIN 
        inventory i ON ts.inv_id = i.id
    JOIN 
        skins s ON i.skin_id = s.id
    GROUP BY 
        td.tradeup_id, td.rarity, td.status, td.total_value
    ORDER BY 
        td.tradeup_id DESC
    LIMIT 5
    `
    rows, err := p.conn.Query(context.Background(), q, userId)
    if err != nil {
        return recentTradeups, err
    }
    defer rows.Close()

    for rows.Next() {
        var tradeup tradeups.Tradeup
        var skinsJSON []byte
        
        err := rows.Scan(&tradeup.Id, &tradeup.Rarity, &tradeup.Status, 
                        &tradeup.Value, &skinsJSON)

        if err != nil {
            return recentTradeups, err
        }

        var tradeupSkins []skins.TradeupSkin
        if err := json.Unmarshal(skinsJSON, &tradeupSkins); err != nil {
            return recentTradeups, err
        }

        tradeup.Skins = tradeupSkins
        recentTradeups = append(recentTradeups, tradeup)
    }

    return recentTradeups, rows.Err()
}
