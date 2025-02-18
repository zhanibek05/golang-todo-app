package todo_app

type User struct {
	id       int    `json:"id"`
	name     string `json:"name"`
	username string `json:"username"`
	password string `json:"password"`
}
