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

func (p *PostgresDB) CreateUser(u *user.User) error {
	q := "insert into users(id, username, email, hash, created_at) values($1,$2,$3,$4,$5)"
	_, err := p.conn.Exec(context.Background(), q, u.Uuid, u.Username, u.Email, u.Hash, time.Now())
	return err
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
    select i.id, i.wear_str, i.wear_num, i.price, i.is_stattrak, to_char(i.created_at, 'YYYY/MM/DD HH12:MI:SS'),
		s.name, s.rarity, s.collection, s.image_key,
        count(*) over (partition by s.image_key) as image_group_count
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
        var imageGroupCount int

		err := rows.Scan(&s.Id, &s.Wear, &s.SkinFloat, &s.Price,
                        &s.IsStatTrak, &s.CreatedAt, &s.Name, &s.Rarity, &s.Collection,
                        &imageKey, &imageGroupCount)
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
				'price', i.price,
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

func AddSkin(s *skins.Skin) error {

	return nil
}
