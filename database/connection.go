package database

import (
	"database/sql"
	"fmt"
	"golang-todolist/helper"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func loadEnv() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

}

func init() {

	loadEnv()

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s",
		helper.GetEnv("DB_USER"),
		helper.GetEnv("DB_PASSWORD"),
		helper.GetEnv("DB_NAME"),
	)

	DB, _ = sql.Open("postgres", connStr)

	if err := DB.Ping(); err != nil {
		panic("Error connecting to database")
	}

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxLifetime(60 * time.Minute)
	DB.SetConnMaxIdleTime(10 * time.Minute)
}
