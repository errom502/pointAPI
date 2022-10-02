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
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}
	if b.Token == "" {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	var ctx context.Context = r.Context()
	if b.Info == "" {
		b.Info = "-"
	}
	if err := sqldb.QueryRow(ctx, `
		select id_user from token where token = $1
	`, b.Token).Scan(&b.Owner); err != nil {
	}
	if b.Owner == "" {
		fmt.Fprintf(w, "bad token")
		return
	}
	if err := sqldb.QueryRow(ctx, `
		insert into Bookmark (name, latitude, longitude, info, owner) VALUES ($1, $2, $3, $4, $5) RETURNING id
	`, b.Name, b.Latitude, b.Longitude, b.Info, b.Owner).Scan(&b.ID); err != nil {
	}
	if err != nil {
		fmt.Fprintf(w, "Bookmark adding went wrong!")
		panic(err)
	} else {
		fmt.Fprintf(w, "Bookmark was successfully added!\nBookmark's id is:%d", b.ID)
	}
}

//encore:api public method=GET path=/bookmarks/:token
func getBookmarks(ctx context.Context, token string) (*models.ListResponse, error) {
	var b models.Bookmarks
	b.Token = token
	if err := sqldb.QueryRow(ctx, `
		select id_user from token where token = $1
	`, b.Token).Scan(&b.Owner); err != nil {
	}
	if b.Owner == "" {
		fmt.Println("bad token")
	} else {
		rows, err := sqldb.Query(ctx, `
		select * from Bookmark where owner = $1
	`, b.Owner)
		defer rows.Close()
		fmt.Println(rows)
		bs := []*models.Bookmarks{}
		for rows.Next() {
			var b models.Bookmarks
			if err := rows.Scan(&b.ID, &b.Name, &b.Latitude, &b.Longitude, &b.Info, &b.Owner); err != nil {
				return nil, err
			}
			bs = append(bs, &b)
		}

		return &models.ListResponse{Bookmarks: bs}, err
	}
	return nil, nil
}

//encore:api public raw method=PATCH path=/bookmark/edit
func editBookmark(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}
	var ctx context.Context = r.Context()
	if b.Token == "" {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	if len(b.Token) != 43 {
		fmt.Fprintf(w, "Bad token")
		return
	}
	var check bool
	if err := sqldb.QueryRow(ctx, `
		select exists(select 1 from bookmark where id = $1)
	`, b.ID).Scan(&check); err != nil {
	}
	fmt.Println("row: ", check)
	if check == false {
		fmt.Fprintf(w, "Bookmark with this ID doesn't exist")
		return
	}
	_, err = sqldb.Query(ctx, `
		update Bookmark set name = $1, latitude = $2, longitude = $3, info = $4 where id = $5
	`, b.Name, b.Latitude, b.Longitude, b.Info, b.ID)
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(w, "Bookmark was successfully updated!")
	}
}

//encore:api public raw method=DELETE path=/bookmark/delete
func deleteBookmark(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}
	var ctx context.Context = r.Context()
	if b.Token == "" {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	if len(b.Token) != 43 {
		fmt.Fprintf(w, "Bad token")
		return
	}
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
