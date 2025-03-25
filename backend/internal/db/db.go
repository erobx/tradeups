package db

import (
	"context"
	"fmt"
	"os"

	"github.com/erobx/tradeups/backend/pkg/url"
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
