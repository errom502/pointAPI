package controllers

import (
	"context"
	"encoding/json"
	"encore.app/models"
	"encore.dev/storage/sqldb"
	"fmt"
	"log"
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

	var check bool
	if err := sqldb.QueryRow(ctx, `
		select exists(select 1 from client where login = $1)
	`, c.Login).Scan(&check); err != nil {
	}
	fmt.Println(err)
	fmt.Println("row: ", check)

	if check == true {
		fmt.Fprintf(w, "This login is already in use")
		return
	}
	log.Println(c)

	_, err = sqldb.Exec(ctx, `
		insert into client (login,password) select $1, $2 where not exists (select login from client where login = $3)
	`, c.Login, c.Password, c.Login)
	if err != nil {
		fmt.Println("Ошибка на 48 строка")
		panic(err)
	}

	fmt.Fprintf(w, "Successful registration")
	fmt.Println("Successful registration")
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
		idFromDB       int
	)

	if err := sqldb.QueryRow(ctx, `
		select id, password from client where login = $1
	`, c.Login).Scan(&idFromDB, &passwordFromDB); err != nil {
	}
	fmt.Println("id from bd: ", idFromDB)
	models.GlobId = idFromDB
	fmt.Println("passw from bd: ", passwordFromDB)
	if c.Password == passwordFromDB {
		fmt.Fprintf(w, "correct data")
		return
	}
	fmt.Fprintf(w, "invalid data")
}

//encore:api public raw method=DELETE path=/client/delete
func deleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, err := sqldb.Exec(ctx, `
		delete from client where id = $1
	`, models.GlobId)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Account deleted")
}
