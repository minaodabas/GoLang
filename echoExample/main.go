package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main(){
	fmt.Println()
	e := echo.New()
	e.GET("/",func (c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.Start(":8000")
}

