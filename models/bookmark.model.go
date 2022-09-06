package models

type Bookmarks struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Owner int
	Info string `json:"info"`
}