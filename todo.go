package todo_app

type TodoList struct {
	id          int    `json:"id"`
	title       string `json:"title"`
	description string `json:"description"`
}

type UsersList struct {
	id     int
	userId int
	listId int
}

type TodoItem struct {
	id          int    `json:"id"`
	title       string `json:"title"`
	description string `json:"description"`
	done        bool   `json:"done"`
}

type ListsItem struct {
	id     int
	listId int
	itemId int
}
