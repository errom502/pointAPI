package client

import (
	"database/sql"
	"encore.app/models"
	"fmt"
)

func Login(client models.Client, db *sql.DB) (bool, string) {
	login := client.Login
	password := client.Password
	passwordFromDB := ""

	rows, err := db.Query(fmt.Sprintf("select password from client where login = '%s'", login))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		erf := rows.Scan(&passwordFromDB)
		if erf != nil {
			panic(err)
		}
		//fmt.Println(resStr)
	}
	if password == passwordFromDB {
		return true, "correct password"
	}
	return false, "invalid data"
}
