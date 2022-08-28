package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	models "encore.app/models"
	"encore.dev/storage/sqldb"
)

//encore:api public raw
func send(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
  err := decoder.Decode(&b)
	if err != nil {
    panic(err)
  }

	log.Println(b)
	var ctx context.Context = r.Context()
	_, err = sqldb.Exec(ctx, `
		INSERT INTO Bookmarks (id, name, address, owner, info) VALUES ($1, $2, $3, $4, $5)
	`, b.ID, b.Name, b.Address, b.Owner, b.Info)
	if err != nil {
    panic(err)
  }
}

//encore:api public method=GET path=/bookmarks
func getBookmarks(ctx context.Context) (*models.ListResponse, error) {
	rows, err := sqldb.Query(ctx, `
		select * from Bookmarks
	`,)
	defer rows.Close()
	bs := []*models.Bookmarks{}

	for rows.Next() {
		var b models.Bookmarks

		if err := rows.Scan(&b.ID, &b.Name, &b.Address, &b.Owner, &b.Info); err != nil {
      return nil, err
    }
    bs = append(bs, &b)
	}

	return &models.ListResponse{Bookmarks: bs}, err
}