package main

import (
	"rax/ShowRater/internal/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func registerGetHandler(c echo.Context) error{
	if !(c.Request().Header.Get("HX-Request") == "true"){
		return c.String(403,"403 Forbidden")
	}
	return c.Render(200,"registerModal", nil)
}

func indexGetHandler(c echo.Context) error{
	return c.Render(200, "index", nil)
}

func usernamePutHandler(c echo.Context) error{
	if !(c.Request().Header.Get("HX-Request") == "true"){
		return c.String(403,"403 Forbidden")
	}
	username := c.FormValue("username")
	if !(len(username) > 0){
		return c.NoContent(200)
	}

	return c.NoContent(400)
}

func setupRoutes(e *echo.Echo){
	e.File("/favicon.ico","images/favicon.ico")
	e.Static("/css", "css")

	e.GET("/", indexGetHandler)
	e.GET("/registerModal", registerGetHandler)
	e.POST("/registerModal/username", usernamePutHandler)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = templates.NewTemplate()

	setupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
