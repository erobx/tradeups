package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/erobx/tradeups/backend/pkg/tradeups"
	"github.com/erobx/tradeups/backend/pkg/url"
	"github.com/erobx/tradeups/backend/pkg/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	conn *pgxpool.Pool
    urlManager *url.PresignedUrlManager
}

func NewPostgresDB() (*PostgresDB, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	bucketName := os.Getenv("S3_BUCKET")
    pm := url.NewPresignedUrlManager(bucketName)

	return &PostgresDB{
        conn: conn,
        urlManager: pm,
    }, nil
}

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
    to_char(created_at, 'YYYY/MM/DD HH12:MI:SS')
    `
	row := p.conn.QueryRow(context.Background(), q, u.Uuid, u.Username, u.Email, u.Hash, time.Now())
    err := row.Scan(&userData.Id, &userData.Username, &userData.Email, &avatarKey, &userData.RefreshTokenVersion, &userData.CreatedAt)

    if avatarKey != "" {
        urlMap := p.urlManager.GetUrls([]string{avatarKey})
        if key, ok := urlMap[avatarKey]; ok {
            avatarKey = key
        }
    }
    userData.AvatarSrc = avatarKey

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
    to_char(created_at, 'YYYY/MM/DD HH12:MI:SS') from users where
    id = $1
    `
    row := p.conn.QueryRow(context.Background(), q, id)
    err := row.Scan(&userData.Id, &userData.Username, &userData.Email, &avatarKey, &userData.RefreshTokenVersion, &userData.CreatedAt)

    if avatarKey != "" {
        urlMap := p.urlManager.GetUrls([]string{avatarKey})
        if key, ok := urlMap[avatarKey]; ok {
            avatarKey = key
        }
    }
    userData.AvatarSrc = avatarKey

	return userData, err
}

// {id: 0, name: "M4A4 | Howl", wear: "Factory New", rarity: "Contraband", float: 0.01, isStatTrak: true, imgSrc: "/m4a4-howl.png"},
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
    select i.id, i.wear_str, i.wear_num, round(cast(i.price as numeric),2), i.is_stattrak, to_char(i.created_at, 'YYYY/MM/DD HH12:MI:SS'),
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

func (p *PostgresDB) IsUsersSkin(userId string, invId int) bool {
    var exists bool
    q := "select exists(select 1 from inventory where user_id=$1 and id=$2)"
    row := p.conn.QueryRow(context.Background(), q, userId, invId)
    row.Scan(&exists)
    return exists
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
