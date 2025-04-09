package router

import (
	"back/Handler"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func RoutersGroup(r *gin.Engine, db *sql.DB) {
	routerGroupTodos := r.Group("/todos")
	{
		routerGroupTodos.GET("/GetAllTodos", Handler.GetHandler(db))
	}
}
