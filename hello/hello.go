package hello

import (
	"context"
	models "encore.app/models"
)

//encore:api public path=/world/:name
func World(ctx context.Context, name string) (*models.Response, error) {
	msg := "world, " + name + "!"
	return &models.Response{Message: msg}, nil
}

//encore:api public path=/hello/:name
func Hello(ctx context.Context, name string) (*models.Response, error) {
	msg := "Hello, " + name + "!"
	return &models.Response{Message: msg}, nil

}
