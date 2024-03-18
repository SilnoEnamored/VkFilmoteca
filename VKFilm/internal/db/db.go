package db

import (
	"github.com/go-pg/pg/v10"
	"os"
)

// Установить значение
func New() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
	return db
}
