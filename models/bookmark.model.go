package models

type Bookmarks struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Address string `json:"addr"`
	Owner string `json:"owner"`
	Info string `json:"info"`
}