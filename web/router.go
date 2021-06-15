package web

import (
	"okakyo/practice-go/config/db"
	"okakyo/practice-go/web/handler"

	"github.com/labstack/echo"
)

func TodoRouting(e *echo.Echo) {
	db,err:= db.NewSqlHandler()
	if err!= nil {
		panic(err)
	}
	router := e.Group("/todos")
	todoHandler :=handler.NewTodoHandler(db)
	router.GET("",todoHandler.GetTodos)
	router.GET("/:id",todoHandler.GetTodo)
	router.POST("",todoHandler.PostTodo)
	router.PUT("/:id",todoHandler.PutTodo)
	router.DELETE("/:id",todoHandler.DeleteTodo)
}

