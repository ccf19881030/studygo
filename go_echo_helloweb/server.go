package main

import (
	api "go_echo_helloweb/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":8921"))
	
}
