package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DatabaseConnect() *sql.DB {
	connStr := "user=postgres dbname=pointapi password=1111 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return db
}

func DatabaseClose(db *sql.DB) {

	defer db.Close()
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

func DatabaseInsert(insert string, db *sql.DB) {
	var err error
	_, err = db.Exec(insert)
	if err != nil {
		panic(err)
	}

}
