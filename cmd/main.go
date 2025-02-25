package main

import (
	"github.com/spf13/viper"
	"github.com/zhanibek05/golang-todo-app"
	"github.com/zhanibek05/golang-todo-app/pkg/database"
	"github.com/zhanibek05/golang-todo-app/pkg/handler"
	"github.com/zhanibek05/golang-todo-app/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}
	db := database.NewDatabase()
	services := service.NewService(db)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("Server Error", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
