package models

type Bookmarks struct {
	Token     string  `json:"token"`
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Info      string  `json:"info"`
	Owner     string
}
