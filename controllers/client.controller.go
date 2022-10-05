package controllers

import (
	"context"
	"encoding/json"
	"encore.app/models"
	"encore.dev/storage/sqldb"
	"fmt"
	"net/http"
)

//encore:api public raw method=POST path=/client/reg
func ClientRegistration(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c models.Client
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	var ctx context.Context = r.Context()

	if c.Login == "" {
		fmt.Fprintf(w, "Send me some normal login,ok?")
		return
	}
	if c.Password == "" {
		fmt.Fprintf(w, "Great password, but can you send me some normal password?")
		return
	}
	var check bool
	if err := sqldb.QueryRow(ctx, `
		select exists(select 1 from client where login = $1)
	`, c.Login).Scan(&check); err != nil {
	}
	fmt.Println("row: ", check)

	if check == true {
		fmt.Fprintf(w, "This login is already in use")
		return
	}

	_, err = sqldb.Exec(ctx, `
		insert into client (login, password) select $1, $2 where not exists(select login from client where login = $3)
	`, c.Login, c.Password, c.Login)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Successful registration")
}

//encore:api public raw method=POST path=/client/login
func ClientLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c models.Client
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	ctx := r.Context()

	var (
		passwordFromDB string
		idFromDB       string
	)
	if c.Login == "" {
		fmt.Fprintf(w, "Send me some normal login,ok?")
		return
	}
	if err := sqldb.QueryRow(ctx, `
		select id, password from client where login = $1
	`, c.Login).Scan(&idFromDB, &passwordFromDB); err != nil {
	}
	if c.Password == passwordFromDB {
		c.Token, _ = GenerateToken(ctx, c)
		fmt.Fprintf(w, "token:%s", c.Token)
		return
	}
	fmt.Fprintf(w, "Something went wrong")
}

//encore:api public raw method=DELETE path=/client/delete
func deleteAccount(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c models.Client
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	ctx := r.Context()
	if len(c.Token) != 43 {
		fmt.Fprintf(w, "Bad token")
		return
	}
	if err := sqldb.QueryRow(ctx, `
		select id_user from token where token = $1
	`, c.Token).Scan(&c.Id); err != nil {
	}
	fmt.Println(c.Token, c.Id)
	_, err = sqldb.Exec(ctx, `
		delete from client where id = $1
	`, c.Id)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Account deleted")
}
