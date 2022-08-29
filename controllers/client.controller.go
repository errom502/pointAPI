package controllers

import (
	"encoding/json"
	"encore.app/models"
	"encore.dev/storage/sqldb"
	"fmt"
	"log"
	"net/http"
)

//encore:api public raw method=POST path=/client/registration
func ClientRegistration(w http.ResponseWriter, r *http.Request) {
	var c models.Client
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}

	ctx := r.Context()
	rows, err := sqldb.Query(ctx, `
		select exists(select 1 from client where login = '$1')
	`, c.Login)
	defer rows.Close()
	check := ""
	for rows.Next() {
		if err := rows.Scan(&check); err != nil {
			panic(err)
		}
		fmt.Println(check)
	}
	if check == "true" {

		fmt.Fprintf(w, "This login is already in use")
	}
	log.Println(c)

	_, err = sqldb.Exec(ctx, `
		INSERT INTO client (login,password) SELECT '$1', '$2' WHERE NOT EXISTS (SELECT login FROM client WHERE login = '$1')
	`, c.Login, c.Password)
	if err != nil {
		panic(err)
	}
	//return "successful registration"
	//return
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
	passwordFromDB := ""
	rows, err := sqldb.Query(ctx, `
		select password from client where login = '$1'
	`, c.Login)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&passwordFromDB); err != nil {
			panic(err)
		}
	}
	if c.Password == passwordFromDB {
		fmt.Fprintf(w, "correct password")
	}
	fmt.Fprintf(w, "invalid data")
}

//DELETE FROM Bookmarks WHERE owner = $1
//DELETE FROM client WHERE id = $1
