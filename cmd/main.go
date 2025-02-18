package main

import (
	todo_app "github.com/zhanibek05/golang-todo-app"
	"github.com/zhanibek05/golang-todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("Server Error", err.Error())
	}

}
