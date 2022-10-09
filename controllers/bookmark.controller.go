package controllers

import (
	_ "encore.dev/middleware"
	"encore.app/models"
	"encore.dev/storage/sqldb"

	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

//encore:api public raw method=POST path=/bookmark/add
func addBookmark(w http.ResponseWriter, r *http.Request) {
	tkn := r.Header.Get("token")
	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	if err != nil {
		panic(err)
	}
	if tkn == "" {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	var ctx context.Context = r.Context()
	if b.Info == "" {
		b.Info = "-"
	}
	
	if err := sqldb.QueryRow(ctx, `
	select id_user from token where token = $1
	`, tkn).Scan(&b.Owner); err != nil {
	}
	if b.Owner == 0 {
		fmt.Fprintf(w, "bad token")
		return
	}
	fmt.Println(&b)
	fmt.Println(b.Owner)
	
	if err := sqldb.QueryRow(ctx, `insert into Bookmark (id, name, latitude, longitude, info, owner) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, models.Concat(b.Owner, ctx), b.Name, b.Latitude, b.Longitude, b.Info, b.Owner).Scan(&b.ID); err != nil {
		fmt.Fprintf(w, "Bookmark adding went wrong!")
		panic(err)
	} else {
		fmt.Fprintf(w, "Bookmark was successfully added!\nBookmark's id is: %d", b.ID)
	}
}

//доделать расшифровку токена

//encore:api public method=GET path=/bookmarks
func getBookmarks(ctx context.Context, m models.Response) (*models.ListResponse, error) {
	tkn := m.Token
	// var ctx context.Context = r.Context()

	var b models.Bookmarks
	if err := sqldb.QueryRow(ctx, `
		select id_user from token where token = $1
	`, tkn).Scan(&b.Owner); err != nil {
	}
	if b.Owner == 0 {
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
	tkn := r.Header.Get("token")

	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}
	var ctx context.Context = r.Context()
	if tkn == "" {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	// if len(b.Token) != 43 {
	// 	fmt.Fprintf(w, "Bad token")
	// 	return
	// }
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
	tkn := r.Header.Get("token")
	var ctx context.Context = r.Context()

	decoder := json.NewDecoder(r.Body)
	var b models.Bookmarks
	err := decoder.Decode(&b)
	fmt.Println(b)
	if err != nil {
		panic(err)
	}
	if tkn == "" {
		fmt.Fprintf(w, "You must be logged in")
		return
	}
	// if len(b.Token) != 43 {
	// 	fmt.Fprintf(w, "Bad token")
	// 	return
	// }
	var check bool
	_ = sqldb.QueryRow(ctx, `
		select exists(select 1 from bookmark where id = $1)
	`, b.ID).Scan(&check)
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

// //encore:middleware global target=tag:getb
// func ValidationMiddleware(req middleware.Request, next middleware.Next) middleware.Response {
// 	// If the payload has a Validate method, use it to validate the request.
// 	payload := req.Data().Payload
// 	print(payload)


// 	// if validator, ok := payload.(interface { Validate() error }); ok {
// 	// 		if err := validator.Validate(); err != nil {
// 	// 				// If the validation fails, return an InvalidArgument error.
// 	// 				err = errs.WrapCode(err, errs.InvalidArgument, "validation failed")
// 	// 				return middleware.Response{Err: err}
// 	// 		}
// 	// }
// 	return next(req)
// }