package database

import (
	"database/sql"
	"fmt"
)

type DBs struct {
	db  *sql.DB
	err error
}

func databaseConnect() {
	var p DBs
	connStr := "user=postgres dbname=pointapi password=1111 host=localhost sslmode=disable"
	p.db, p.err = sql.Open("postgres", connStr)
	fmt.Printf("\nSuccessfully connected to database!\n")

}

func databaseClose() {
	var p DBs
	defer p.db.Close()
	p.err = p.db.Ping()
	if p.err != nil {
		panic(p.err)
	}
}
