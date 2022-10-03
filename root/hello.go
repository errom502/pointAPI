package hello

import (
	"context"
	"encore.app/models"
)

//	msg := "Welcome to our API!\nHere's examples how you can use api:\n/client/reg\n{\n    \"login\": \"your_login\",\n    \"password\": \"password\"\n}\n/client/login\n{\n    \"login\": \"your_login\",\n    \"password\": \"password\"\n}\n/bookmark/add\n{\n    \"token\": \"your_token\",\n    \"name\": \"name of your bookmark\",\n    \"latitude\": 34.2242,\n    \"longitude\": 7.2325,\n    \"info\": \"some info of bookmark\"\n}\n/bookmarks/your_token\n\n/bookmark/edit\n{\n    \"id\": 0,\n    \"token\": \"your_token\",\n    \"name\": \"CHANGED NAME\",\n    \"latitude\": 34.2242,\n    \"longitude\": 7.2325,\n    \"info\": \"some info of bookmark\"\n}\n/client/delete\n{\n    \"token\": \"your token\"\n}\n/bookmark/delete\n{\n    \"token\": \"your_token\",\n    \"id\": 1\n}"
//
//encore:api public path=/hello
func HelloWorld(ctx context.Context) (*models.Response, error) {
	msg := "Welcome to our API!Here's examples how you can use api:\n/client/reg\n{\n    \"login\": \"your_login\",\n    \"password\": \"password\"\n}\n/client/login\n{\n    \"login\": \"your_login\",\n    \"password\": \"password\"\n}\n/bookmark/add\n{\n    \"token\": \"your_token\",\n    \"name\": \"name of your bookmark\",\n    \"latitude\": 34.2242,\n    \"longitude\": 7.2325,\n    \"info\": \"some info of bookmark\"\n}\n/bookmarks/your_token\n\n/bookmark/edit\n{\n    \"id\": 0,\n    \"token\": \"your_token\",\n    \"name\": \"CHANGED NAME\",\n    \"latitude\": 34.2242,\n    \"longitude\": 7.2325,\n    \"info\": \"some info of bookmark\"\n}\n/client/delete\n{\n    \"token\": \"your token\"\n}\n/bookmark/delete\n{\n    \"token\": \"your_token\",\n    \"id\": 1\n}"
	return &models.Response{Message: msg}, nil
}
