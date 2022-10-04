package controllers

import (
	"context"
	"encore.dev/storage/sqldb"
)

// create token after login
// select * from token ORDER BY date_create desc limit 5;

func CreateToken(ctx context.Context, id string) (token string) {
	_, err := sqldb.Exec(ctx, `
		insert into token(id_user) values($1);
	`, id)
	if err != nil {
		panic(err)
	}
	if err := sqldb.QueryRow(ctx, `
		select token from token where id_user = $1 order by date_create desc limit 1
	`, id).Scan(&token); err != nil {
	}
	if err != nil {
		panic(err)
	}
	if token == "" {
		return "Error token create"
	}
	return token
}
