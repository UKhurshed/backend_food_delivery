package main

import "backend_food_delivery/internal/app"

const  configPath = "configs/main"

func main()  {

	app.Run(configPath)
	//handlers := new(handler.Handler)
	//srv := new(server.Server)
	//if err := srv.Run("8080", handlers.InitRoutes()); err!= nil{
	//	log.Fatalf("error occurred while running http server: %s", err.Error())
	//}
}