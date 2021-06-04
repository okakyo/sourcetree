package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

func TodoRouting(e *echo.Echo) {
	router := e.Group("/users")
	router.GET("",func(c echo.Context) error {
		return c.JSON(http.StatusOK,map[string]string {"hello":"World"})
	})
	router.GET("/:id",func(c echo.Context) error {
		return c.JSON(http.StatusOK,map[string]string{"good":"morning"})
	})
	router.POST("",func (c echo.Context) error {
		return c.JSON(http.StatusCreated,map[string]string {"Good":"Afternoon"})
	})

	router.PUT("/:id",func(c echo.Context) error {
		return c.JSON(http.StatusOK,map[string]string {"Good":"Evening"})
	})

	router.DELETE("/:id",func(c echo.Context) error {
		return c.JSON(http.StatusOK,map[string]string {"Good":"Night"})
	})
}

