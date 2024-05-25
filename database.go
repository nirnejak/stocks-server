package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func LoadEnv() error {
	return godotenv.Load(".env")
}

func InitDB() error {
	var err error
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		return err
	}
	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
