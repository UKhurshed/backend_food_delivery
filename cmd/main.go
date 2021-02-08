package main

import (
	"backend_food_delivery/pkg/handler"
	"backend_food_delivery/server"
	"log"
)

func main()  {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err!= nil{
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}