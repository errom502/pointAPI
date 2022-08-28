package main

import (
	"encore.app/client"
	"encore.app/database"
	"encore.app/models"
	"fmt"
)

func main() {
	db := database.DatabaseConnect()
	var c models.Client
	c.Login = "Megat fj1df"
	c.Password = "ssega 2"
	//reg_result := client.Registration(c, db) //проверяет используется логин или нет
	//fmt.Println(reg_result)
	fmt.Println(client.Login(c, db))
}
