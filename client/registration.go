package client

import (
	"database/sql"
	"encore.app/database"
	"encore.app/models"
	"fmt"
)

//    insert into client (login, password) values ('Loginmega1','sdkwdk2');
func Registration(client models.Client, db *sql.DB) {
	login := client.Login
	password := client.Password
	str := "INSERT INTO client (login,password) VALUES ('" + login + "','" + password + "');"
	fmt.Println(str)
	database.DatabaseInsert(str, db)
}
