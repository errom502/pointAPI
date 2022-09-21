package controllers

import (
	"context"
	"encoding/json"
	"encore.app/models"
	"encore.dev/storage/sqldb"
	"fmt"
	"net/http"
)

//encore:api public raw method=POST path=/bookmark/add
func addBookmark(w http.ResponseWriter, r *http.Request) {
	if models.GlobId == 0 {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}

	var ctx context.Context = r.Context()
	if b.Info == "" {
		b.Info = "-"
	}
	_, err = sqldb.Exec(ctx, `
		insert into Bookmark (name, address, info, coordinates, owner) VALUES ($1, $2, $3, $4, $5)
	`, b.Name, b.Address, b.Info, b.Coordinates, models.GlobId)
	if err != nil {
		fmt.Fprintf(w, "Bookmark adding went wrong!")
		panic(err)
	} else {
		fmt.Fprintf(w, "Bookmark was successfully added!")
	}
}

//encore:api public method=GET path=/bookmarks
func getBookmarks(ctx context.Context) (*models.ListResponse, error) {
	rows, err := sqldb.Query(ctx, `
		select * from Bookmark where owner = $1
	`, models.GlobId)
	defer rows.Close()
	bs := []*models.Bookmarks{}

	for rows.Next() {
		var b models.Bookmarks

		if err := rows.Scan(&b.ID, &b.Name, &b.Address, &b.Info, &b.Coordinates, &b.Owner); err != nil {
			return nil, err
		}
		bs = append(bs, &b)
	}

	return &models.ListResponse{Bookmarks: bs}, err
}

//encore:api public raw method=PATCH path=/bookmark/edit
func editBookmark(w http.ResponseWriter, r *http.Request) {
	if models.GlobId == 0 {
		fmt.Fprintf(w, "You must be logged in")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}
	var ctx context.Context = r.Context()
	row, err := sqldb.Query(ctx, `
		update Bookmark set name = $1, address = $2, info = $3, coordinates = $4 where id = $5
	`, b.Name, b.Address, b.Info, &b.Coordinates, b.ID)
	fmt.Println(row)
	if err != nil {
		fmt.Fprintf(w, "Bookmark with this ID doesn't exist")
		panic(err)
	} else {
		fmt.Fprintf(w, "Bookmark was successfully updated!")
	}
}

//encore:api public raw method=DELETE path=/bookmark/delete
func deleteBookmark(w http.ResponseWriter, r *http.Request) {
	if models.GlobId == 0 {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}
	var ctx context.Context = r.Context()
	var check bool
	if err := sqldb.QueryRow(ctx, `
		select exists(select 1 from bookmark where id = $1)
	`, b.ID).Scan(&check); err != nil {
	}

	if check == false {
		fmt.Fprintf(w, "Bookmark with this ID doesn't exist")
		return
	}

	_, err = sqldb.Exec(ctx, `
		delete from bookmark where id = $1
	`, b.ID)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Bookmark deleted")
}
