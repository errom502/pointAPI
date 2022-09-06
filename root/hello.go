package hello

import (
	"context"
	"encore.app/models"
)

//encore:api public path=/hello
func HelloWorld(ctx context.Context) (*models.Response, error) {
	msg := "Welcome to our API!"
	return &models.Response{Message: msg}, nil
}