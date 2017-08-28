
package main

import (
"net/http"
"html/template"
"github.com/labstack/echo"
	"io"
)
// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func search(c echo.Context) error{

}

func main() {
	e := echo.New()
	e.File("/", "src/client/index.html")
	e.POST("/search", search)
	e.Logger.Fatal(e.Start(":1323"))
}
