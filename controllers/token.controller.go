package controllers

import (
	"context"
	"encore.app/models"
	"encore.dev/storage/sqldb"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// create token after login
// select * from token ORDER BY date_create desc limit 5;

//func CreateToken(ctx context.Context, id string) (token string) {
//	_, err := sqldb.Exec(ctx, `
//		insert into token(id_user) values($1);
//	`, id)
//	if err != nil {
//		panic(err)
//	}
//	if err := sqldb.QueryRow(ctx, `
//		select token from token where id_user = $1 order by date_create desc limit 1
//	`, id).Scan(&token); err != nil {
//	}
//	if err != nil {
//		panic(err)
//	}
//	if token == "" {
//		return "Error token create"
//	}
//	return token
//}
//

const (
	salt       = "ndjfij38j8hf7dgf73"
	signingKey = "2k2je9h8hg83ij34h384h394jh9h94j"
	tokenTTL   = 12 * time.Hour // токен перестанет быть валидным через 12 часов
)

func GenerateToken(ctx context.Context, user models.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // токен перестанет быть валидным через 12 часов
			IssuedAt:  time.Now().Unix(),
		},
		Username: user.Login,
	})
	FnlToken, _ := token.SignedString([]byte(signingKey))
	_, err := sqldb.Exec(ctx, `
			insert into token(id_user,token) values($1,$2);
		`, user.Id, FnlToken)
	if err != nil {
		panic(err)
	}
	fmt.Println(FnlToken, user.Id)
	return FnlToken, err
}
