package main 

import (
	"net/http"
	"github.com/labstack/echo"
	"./controllers"
)

func main(){
	e:= echo.New()
	e.Use()
	e.GET("/",func (c echo.Context) error {
		return c.String(http.StatusOK,"Hello June")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
