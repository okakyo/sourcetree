package main 

import (
	"net/http"
	"github.com/labstack/echo"
	"okakyo/practice-go/web"
	"okakyo/practice-go/database"
)

func main(){
	e:= echo.New()
	web.TodoRouting(e)
	e.GET("/",func (c echo.Context) error {
		return c.String(http.StatusOK,"Hello June")
	})
	e.PUT("/database/reset/all",database.ResetDB)
	e.Logger.Fatal(e.Start(":8000"))
}
