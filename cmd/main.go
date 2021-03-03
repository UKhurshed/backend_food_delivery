package main

import (
	"backend_food_delivery/pkg/handler"
	"backend_food_delivery/pkg/repository"
	"backend_food_delivery/pkg/service"
	"backend_food_delivery/server"
	"fmt"
	"log"
)

func main()  {
	fmt.Print("Hello")
	db, err := repository.NewClient("mongodb://127.0.0.1:27017")
	if err != nil{
		log.Fatalf("Mongo db error: %s",err.Error())
	}

	client := db.Database("testdb")

	repos := repository.NewRepository(client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil{
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
