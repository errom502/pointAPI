package models

type Client struct {
	Login    string `json:"login"`
	Id       int
	Password string `json:"password"`
	Token    string
}

var GlobId int
