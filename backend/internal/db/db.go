package db

import (
	"context"
	"fmt"
	"os"

	"github.com/erobx/tradeups/backend/pkg/lock"
	"github.com/erobx/tradeups/backend/pkg/url"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	conn *pgxpool.Pool
    urlManager *url.PresignedUrlManager
    LockManager *lock.LockManager
}

func NewPostgresDB() (*PostgresDB, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	bucketName := os.Getenv("S3_BUCKET")
    pm := url.NewPresignedUrlManager(bucketName)
    lm := lock.NewLockManager()

	return &PostgresDB{
        conn: conn,
        urlManager: pm,
        LockManager: lm,
    }, nil
}


