package app

import (
	"backend_food_delivery/internal/config"
	"backend_food_delivery/pkg/database/mongodb"
	"backend_food_delivery/pkg/logger"
	"fmt"
)

func Run(configPath string)  {
	cfg, err := config.Init(configPath)
	if err != nil{
		logger.Error(err)
		return
	}

	mongoClient, err := mongodb.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		logger.Error(err)
	}

	db := mongoClient.Database(cfg.Mongo.Name)

	fmt.Printf(db.Name())

	//memCache := cache.NewMemoryCache()




}
