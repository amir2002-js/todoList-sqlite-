package Handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Todos struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// GetHandler GET ALL
func GetHandler(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		row, err := db.Query(`SELECT id ,title FROM "todo"`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		defer func(row *sql.Rows) {
			err := row.Close()
			if err != nil {
				log.Println(err)
			}
		}(row)

		var todos []Todos
		for row.Next() {
			var t Todos
			err = row.Scan(&t.ID, &t.Title)
			if err != nil {
				panic(err)
			}
			todos = append(todos, t)
		}

		c.JSON(http.StatusOK, todos)
	}
}

// GetHandlerById Get by ID
func GetHandlerById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		parId := c.Param("id")

		data := db.QueryRow("SELECT id , title FROM todo WHERE id = ?", parId)

		var todo Todos

		err := data.Scan(&todo.ID, &todo.Title)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusFound, gin.H{"message": "not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

// PostHandler post
func PostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//codes
	}
}

// PutHandler put
func PutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//codes
	}
}

// DeleteHandler delete
func DeleteHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//codes
	}
}
