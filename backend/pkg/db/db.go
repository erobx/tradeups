package db

import (
	"context"
	"fmt"
	"os"

	"github.com/erobx/tradeups/pkg/user"
	"github.com/jackc/pgx/v5"
)

type Postgresql struct {
	conn *pgx.Conn
}

func NewPostgresql() *Postgresql {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return &Postgresql{
		conn: conn,
	}
}

func (p *Postgresql) InsertUser(u User) error {
	return nil
}
