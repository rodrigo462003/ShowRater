package templates


import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type templates struct {
	templates *template.Template
}

func (t *templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *templates {
	return &templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

