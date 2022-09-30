package models

type Client struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var GlobId string