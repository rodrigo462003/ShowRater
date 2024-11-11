package main

import (
	"rax/ShowRater/cmd/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = templates.NewTemplate()

	e.GET("/", func (c echo.Context) error {
		return c.Render(200, "index","hello")
	})

	e.Logger.Fatal(e.Start(":8000"))
}

