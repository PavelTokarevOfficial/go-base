package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// test and connect DB
	db, err := connectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// ====== api start ======
	mux := http.NewServeMux()

	// == ping ==
	mux.HandleFunc("/api/v1/ping", pingHandler)

	// == images ==
	mux.HandleFunc("/api/v1/images", handleImages)
	mux.HandleFunc("/api/v1/images/", handleImageByID)

	// == posts ==

	//create
	//update
	//read all
	//read one by id
	//delete

	log.Println("server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
	// ====== api end ======
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
