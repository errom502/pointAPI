package main

import (
	"encore.app/client"
	"encore.app/database"
	"encore.app/models"
)

func main() {
	db := database.DatabaseConnect()
	var c models.Client
	c.Login = "Mega 1"
	c.Password = "ssega 2"
	client.Registration(c, db) // ну это типо тест еще
}
