package main

import (
	"database/sql"
	"rax/ShowRater/internal/database"
	"rax/ShowRater/internal/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func registerGetHandler(c echo.Context) error{
	if !(c.Request().Header.Get("HX-Request") == "true"){
		return c.NoContent(403)
	}
	return c.Render(200,"registerModal", nil)
}

func indexGetHandler(c echo.Context) error{
	return c.Render(200, "index", nil)
}

func usernameCheckHandler(c echo.Context, db *sql.DB) (err error){
	if !(c.Request().Header.Get("HX-Request") == "true"){
		return c.NoContent(403)
	}

	username := c.FormValue("username")
	uLen := len(username)
	if !(30 >= uLen && uLen >= 4){
		errorMessage := "Error"
		return c.Render(422,"usernameError", errorMessage)
	}

	exists, err := database.UserNameExists(username, db)
	if err != nil{
		c.NoContent(500)
	}
	if exists {
		return c.String(409,"Username already exists")
	}

	return c.NoContent(204)
}

func setupRoutes(e *echo.Echo, db *sql.DB){
	e.File("/favicon.ico","images/favicon.ico")
	e.Static("/css", "css")

	e.GET("/", indexGetHandler)
	e.GET("/registerModal", registerGetHandler)
	e.POST("/registerModal/username", func(c echo.Context) error{
		return usernameCheckHandler(c, db)
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = templates.NewTemplate()

	connectionString := "postgres://admin:admin@db:5432/ShowRater?sslmode=disable"

	db, err := database.GetConnectionPool(connectionString)
	if(err != nil){
		e.Logger.Fatal(err.Error())
	}

	setupRoutes(e,db)

	e.Logger.Fatal(e.Start(":8000"))
}
