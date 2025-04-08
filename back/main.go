package main

import (
	"back/Handler"
	"back/createTable"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./DB/todo.db")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	createTable.CreateTable(db)

	r.GET("/", Handler.GetHandler(db))

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
