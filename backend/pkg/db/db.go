package db

import (
	"context"
	"fmt"
	"os"

	"github.com/erobx/tradeups/backend/pkg/skins"
	"github.com/erobx/tradeups/backend/pkg/user"
	"github.com/jackc/pgx/v5"
)

var Postgresql *pgx.Conn

func Connect() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Postgresql = conn
}

func CreateUser(u *user.User) error {
	q := "insert into users(id, username, email, hash) values($1,$2,$3,$4)"
	_, err := Postgresql.Exec(context.Background(), q, u.Uuid, u.Username, u.Email, u.Hash)
	return err
}

func AddSkin(s *skins.Skin) error {

	return nil
}
