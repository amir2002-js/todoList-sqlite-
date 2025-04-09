package router

import (
	"back/Handler"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func RoutersGroup(r *gin.Engine, db *sql.DB) {
	routerGroupTodos := r.Group("/todos")
	{
		routerGroupTodos.GET("/", Handler.GetHandler(db))
		routerGroupTodos.GET("/:id", Handler.GetHandlerById(db))
		routerGroupTodos.DELETE("/:id", Handler.DeleteHandler(db))
		routerGroupTodos.POST("/", Handler.PostHandler(db))
		routerGroupTodos.PUT("/:id", Handler.PutHandler(db))
	}
}
