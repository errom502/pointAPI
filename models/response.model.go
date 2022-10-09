package models

type Response struct {
	Message string ''
	Token   string `header:"token"`
}