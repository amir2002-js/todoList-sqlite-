package main

import (
	"back/Handler"
	"back/createTable"
	"database/sql"
	_ "github.com/glebarez/sqlite"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite", "./DB/todo.db")
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	createTable.CreateTable(db)

	http.HandleFunc("/", Handler.Handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
