package Handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todos struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
	About  string `json:"about"`
}

// GetHandler GET ALL
func GetHandler(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		row, err := db.Query(`SELECT id ,title , about , isDone FROM "todo"`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		defer func(row *sql.Rows) {
			err := row.Close()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}(row)

		var todos []Todos
		for row.Next() {
			var t Todos
			err = row.Scan(&t.ID, &t.Title, &t.About, &t.IsDone)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			todos = append(todos, t)
		}

		c.JSON(http.StatusOK, todos)
	}
} // ✅

// GetHandlerById Get by ID
func GetHandlerById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		parId := c.Param("id")

		data := db.QueryRow("SELECT id , title , about , isDone FROM todo WHERE id = ?", parId)

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
} // ✅

// PostHandler post
func PostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo Todos
		err := c.ShouldBindJSON(&todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, err := db.Exec("INSERT INTO todos (title , about , isDone) VALUES (? , ? , ?)", todo.Title, todo.About, todo.IsDone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ourId, _ := result.LastInsertId()
		todo.ID = int(ourId)

		c.JSON(http.StatusCreated, todo)

	}
} // ✅

// PutHandler update
func PutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo Todos
		parId := c.Param("id")

		err := c.ShouldBindJSON(&todo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, err := db.Exec("UPDATE todo SET title = ?, about = ?, isDone = ? WHERE id = ?", todo.Title, todo.About, todo.IsDone, parId)
		if index, err := result.LastInsertId(); index == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
	}
}

// DeleteHandler delete
func DeleteHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		parId := c.Param("id")

		result, err := db.Exec("DELETE FROM todo WHERE id = ?", parId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
	}
} // ✅
