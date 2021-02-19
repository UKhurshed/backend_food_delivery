package main

import (
	"backend_food_delivery/pkg/handler"
	"backend_food_delivery/pkg/repository"
	"backend_food_delivery/pkg/service"
	"backend_food_delivery/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs :%s", err.Error())
	}

	if err := godotenv.Load("../.env"); err != nil{
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db := repository.InitDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
	//var ctx = context.Context()

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
