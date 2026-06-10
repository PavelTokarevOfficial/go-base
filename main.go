package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	dsn := "postgres://go_user:go_pass@127.0.0.1:5433/go_db?sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {

		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected")
}
