package client

import (
	"database/sql"
	"encore.app/database"
	"encore.app/models"
	"fmt"
)

//    INSERT INTO users (id,username) SELECT %d, '%s' WHERE NOT EXISTS (SELECT id FROM users WHERE id = %d)
// 	  INSERT INTO client (login,password) SELECT %d, '%s' WHERE NOT EXISTS (SELECT id FROM users WHERE id = %d)

func Registration(client models.Client, db *sql.DB) string {
	login := client.Login
	password := client.Password

	existStr := "select exists(select 1 from client where login = '" + login + "');"
	rows, err := db.Query(existStr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	resStr := ""
	for rows.Next() {
		erf := rows.Scan(&resStr)
		if erf != nil {
			panic(err)
		}
		fmt.Println(resStr)
	}
	if resStr == "true" {
		return "This login is already in use"
	}

	str := fmt.Sprintf("INSERT INTO client (login,password) SELECT '%s', '%s' WHERE NOT EXISTS (SELECT login FROM client WHERE login = '%s')", login, password, login)
	fmt.Println(str)
	database.DatabaseInsert(str, db)

	return "successful registration"
}

//select exists(select 1 from client where login = 'Mega fj1');
