package main 

import (
	"net/http"
	"github.com/labstack/echo"
	"okakyo/practice-go/controllers"
)

func main(){
	e:= echo.New()
	controllers.TodoRouting(e)
	e.GET("/",func (c echo.Context) error {
		return c.String(http.StatusOK,"Hello June")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
