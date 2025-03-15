package db

import (
	"context"
	"time"

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

func (p *PostgresDB) IsUsersSkin(userId string, invId int) bool {
    var exists bool
    q := "select exists(select 1 from inventory where user_id=$1 and id=$2)"
    row := p.conn.QueryRow(context.Background(), q, userId, invId)
    row.Scan(&exists)
    return exists
}
