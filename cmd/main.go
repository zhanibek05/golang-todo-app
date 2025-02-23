package main

import (
	"github.com/zhanibek05/golang-todo-app"
	"github.com/zhanibek05/golang-todo-app/pkg/database"
	"github.com/zhanibek05/golang-todo-app/pkg/handler"
	"github.com/zhanibek05/golang-todo-app/pkg/service"
	"log"
)

func main() {
	db := database.NewDatabase()
	services := service.NewService(db)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("Server Error", err.Error())
	}

}
