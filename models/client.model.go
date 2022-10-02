package models

type Client struct {
	Login    string `json:"login"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

var GlobId string
