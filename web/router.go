package controllers

import (
	"github.com/labstack/echo"
	"okakyo/practice-go/usecase"
	"okakyo/practice-go/web/handler"
)

func TodoRouting(e *echo.Echo) {
	router := e.Group("/todos")
	
	router.GET("/:id",usecase.FindOne)
	router.POST("",usecase.CreateTodo)
	router.PUT("/:id",usecase.UpdateTodo)
	router.DELETE("/:id",usecase.DeleteTodo)
}

