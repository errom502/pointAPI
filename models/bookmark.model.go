package models

type Bookmarks struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Info        string `json:"info"`
	Coordinates string `json:"coordinates"`
	Owner       int
}
