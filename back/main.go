package main

import (
	"back/createTable"
	"back/router"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/sqlite"
)

func main() {
	// connect to db
	db, err := sql.Open("sqlite", "./DB/todo.db")
	if err != nil {
		panic(err)
	}

	//close db
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// create an instance
	r := gin.Default()

	//create a table in our db
	createTable.CreateTable(db)

	//group routing
	router.RoutersGroup(r, db)

	//running on 8080 port
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
