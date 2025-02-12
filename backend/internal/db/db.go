package db

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

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

func AddSkin(s *skins.Skin) error {

	return nil
}
