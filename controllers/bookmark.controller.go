package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	models "encore.app/models"
	"encore.dev/storage/sqldb"
)

//encore:api public raw path=/bookmark/add
func addBookmark(w http.ResponseWriter, r *http.Request) {
	if models.GlobId == 0 {
		fmt.Fprintf(w, "You must be logged in")
		return
	} else {
		decoder := json.NewDecoder(r.Body)
		var b models.Bookmarks
		err := decoder.Decode(&b)
		log.Println(b)
		if err != nil {
			panic(err)
		}

		var ctx context.Context = r.Context()
		if b.Info == "" {
			fmt.Println("info is nill")
			b.Info = "-"
		}
		_, err = sqldb.Exec(ctx, `
			insert into Bookmark (name, address, owner, info) VALUES ($1, $2, $3, $4)
		`, b.Name, b.Address, models.GlobId, b.Info)
		if err != nil {
			fmt.Fprintf(w, "Bookmark adding went wrong!")
			panic(err)
		} else {
			fmt.Fprintf(w, "Bookmark was seccessfully added!")
		}
	}
}

//encore:api public method=GET path=/bookmarks
func getBookmarks(ctx context.Context) (*models.ListResponse, error) {
	rows, err := sqldb.Query(ctx, `
		select * from Bookmark
	`,)
	defer rows.Close()
	bs := []*models.Bookmarks{}

	for rows.Next() {
		var b models.Bookmarks

		if err := rows.Scan(&b.ID, &b.Name, &b.Address, &b.Info, &b.Owner); err != nil {
			return nil, err
		}
		bs = append(bs, &b)
	}

	return &models.ListResponse{Bookmarks: bs}, err
}
