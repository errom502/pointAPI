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

//encore:api public raw method=POST path=/client/registration
func ClientRegistration(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c models.Client
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	var ctx context.Context = r.Context()
	//
	var check bool
	if err := sqldb.QueryRow(ctx, `
		select exists(select 1 from client where login = $1)
	`, c.Login).Scan(&check); err != nil {
	}
	fmt.Println(err)
	fmt.Println("row: ", check)
	//
	if check == true {
		fmt.Fprintf(w, "This login is already in use")
		return
	}
	log.Println(c)
	//
	_, err = sqldb.Exec(ctx, `
		insert into client (login,password) select $1, $2 where not exists (select login from client where login = $3)
	`, c.Login, c.Password, c.Login)
	if err != nil {
		fmt.Println("Ошибка на 48 строка")
		panic(err)
	}
	//
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

	var passwordFromDB string
	if err := sqldb.QueryRow(ctx, `
		select password from client where login = $1
	`, c.Login).Scan(&passwordFromDB); err != nil {
	}
	fmt.Println("passw from bd: ", passwordFromDB)
	if c.Password == passwordFromDB {
		fmt.Fprintf(w, "correct data")
		return
	}
	fmt.Fprintf(w, "invalid data")

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
	var id int
	if err := sqldb.QueryRow(ctx, `
		select id from client where login = $1
	`, c.Login).Scan(&id); err != nil {
	}
	fmt.Println(id)

	_, err = sqldb.Exec(ctx, `
		delete from bookmarks where id = $1
	`, id)
	if err != nil {
		panic(err)
	}

	_, err = sqldb.Exec(ctx, `
		delete from client where id = $1
	`, id)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Account deleted")
}

//DELETE FROM Bookmarks WHERE owner = $1
//DELETE FROM client WHERE id = $1
