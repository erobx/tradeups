package db

import (
	"context"
    "encoding/json"
	"fmt"
	"os"

	"github.com/erobx/tradeups/backend/internal/url"
	"github.com/erobx/tradeups/backend/pkg/tradeups"
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

func (p *PostgresDB) getTradeup(tradeupId string) (tradeups.Tradeup, error) {
    var t tradeups.Tradeup
    q := `
    select t.id tradeup_id, t.rarity, t.current_status, t.stop_time,
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
    row := p.conn.QueryRow(context.Background(), q, tradeupId)
    err := row.Scan(&t.Id, &t.Rarity, &t.Status, &t.StopTime, &playersJson, &skinsJson)
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

func (p *PostgresDB) getSkinCount(tradeupId string) (int, error) {
    var numSkins int
    q := "select count(tradeup_id) from tradeups_skins where tradeup_id=$1"
    row := p.conn.QueryRow(context.Background(), q, tradeupId)
    err := row.Scan(&numSkins)
    if err != nil {
        return numSkins, err
    }
    return numSkins, err
}
