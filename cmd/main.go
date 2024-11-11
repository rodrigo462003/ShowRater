package main

import (
	"rax/ShowRater/cmd/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func indexGetHandler(c echo.Context) error{
	return c.Render(200, "index","hello")
}

func setupRoutes(e *echo.Echo){
	e.File("/favicon.ico","images/favicon.ico")
	e.Static("/css", "css")
	e.GET("/", indexGetHandler)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = templates.NewTemplate()

	setupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}

