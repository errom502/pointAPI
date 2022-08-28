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
	c.Login = "Megatrsr fj1df"
	c.Password = "ssega 2"
	reg_result := client.Registration(c, db) //проверяет используется логин или нет
	fmt.Println(reg_result)
}
