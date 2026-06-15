package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("DB connected")
}

func connectDB() (*sql.DB, error) {
	dsn := "postgres://go_user:go_pass@127.0.0.1:5433/go_db?sslmode=disable"

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
