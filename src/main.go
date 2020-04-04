package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
