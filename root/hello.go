package hello

import (
	"context"
	"encore.app/models"
	"fmt"
)

//encore:api public path=/hello
func HelloWorld(ctx context.Context) (*models.Response, error) {
	msg := fmt.Sprintf("Welcome to our API!<br>Here's examples how you can use api:<br>Add your token in header into key \"token\" in every request!<br>REGISTRATION:<br>/client/reg<br>{<br>   \"login\": \"your_login\",<br>    \"password\": \"password\"<br>}<br>LOGIN:<br>/client/login{<br>    \"login\": \"your_login\",<br>    \"password\": \"password\"<br>}<br>ADD BOOKMARK:<br>/bookmark/add<br>{<br>    \"name\": \"name of your bookmark\",<br>    \"latitude\": 34.2242,<br>    \"longitude\": 7.2325,<br>    \"info\": \"some info of bookmark\"<br>}<br>GET YOUR BOOKMARKS:<br>/bookmarks/<br>BOOKMARK EDIT:<br>/bookmark/edit<br>{<br>    \"id\": 0,<br>    \"name\": \"CHANGED NAME\",<br>    \"latitude\": 34.2242,<br>    \"longitude\": 7.2325,<br>    \"info\": \"some info of bookmark\"<br>}<br>DELETE CLIENT:<br>/client/delete<br>DELETE BOOKMARK:<br>/bookmark/delete<br>{<br>    \"id\": 1<br>}<br><br>")
	return &models.Response{Message: msg}, nil
}
