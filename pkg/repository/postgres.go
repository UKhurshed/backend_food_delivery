package repository

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
)

const (
	userTable = "userFood"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func InitDB(cfg Config) *pg.DB {
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	opts := &pg.Options{
		Database: cfg.DBName,
		User:     cfg.Username,
		Password: cfg.Password,
		Addr:     address,
	}

	var db = pg.Connect(opts)

	if db == nil {
		log.Print("Failed to connect")
	}

	return db

}
