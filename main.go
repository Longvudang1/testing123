package main

import (
	"github.com/Longvudang1/testing123/handler"
	"github.com/Longvudang1/testing123/mdw"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())
	server.GET("/", handler.Hello)
	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))
	server.Logger.Fatal(server.Start(":8080"))
}
