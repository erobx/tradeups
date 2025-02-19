package db

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/erobx/tradeups/backend/pkg/common"
	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/erobx/tradeups/backend/pkg/user"
	"github.com/jackc/pgx/v5"
)

type PostgresDB struct {
	mu sync.RWMutex
	conn *pgx.Conn
}

func NewPostgresDB() (*PostgresDB, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return &PostgresDB{conn: conn}, nil
}

func (p *PostgresDB) FindEmail(email string) (bool, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var exists bool
	q := "select exists(select 1 from users where email=$1 limit 1)"
	row := p.conn.QueryRow(context.Background(), q, email)
	row.Scan(&exists)

	return exists, nil
}

func (p *PostgresDB) FindUsername(username string) (bool, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var exists bool
	q := "select exists(select 1 from users where username=$1 limit 1)"
	row := p.conn.QueryRow(context.Background(), q, username)
	row.Scan(&exists)

	return exists, nil
}

func (p *PostgresDB) CreateUser(u *user.User) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	q := "insert into users(id, username, email, hash, created_at) values($1,$2,$3,$4,$5)"
	_, err := p.conn.Exec(context.Background(), q, u.Uuid, u.Username, u.Email, u.Hash, time.Now())
	return err
}

func (p *PostgresDB) GetHash(email string) (id, hash string, err error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

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
	p.mu.RLock()
	defer p.mu.RUnlock()

	var inv user.Inventory
	var items []skins.InventorySkin
	q :=`
	select us.id, us.skin_float, us.skin_price, us.is_stattrak, us.wear,
	s.name, s.rarity, s.collection, s.image_key
	from user_skins us
		join skins s on s.id = us.skin_id
	where us.user_id=$1
	`
	rows, err := p.conn.Query(context.Background(), q, userId)
	if err != nil {
		return inv, err
	}
	defer rows.Close()

	for rows.Next() {
		var s skins.InventorySkin
		var imageKey string

		err := rows.Scan(&s.Id, &s.SkinFloat, &s.SkinPrice, &s.IsStatTrak, &s.Wear, &s.Name, &s.Rarity, &s.Collection, &imageKey)
		if err != nil {
			return inv, err
		}

		imgSrc := common.GetPresignedURL(imageKey)
		s.ImageSrc = imgSrc
		
		items = append(items, s)
	}

	inv.Skins = items
	return inv, rows.Err()
}

func AddSkin(s *skins.Skin) error {

	return nil
}
