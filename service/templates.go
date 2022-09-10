package service

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(templates string) *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(templates)),
	}
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {

	if data == nil {
		data = make(map[string]any)
	}

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]any); isMap {
		viewContext["reverse"] = c.Echo().Reverse
		viewContext["isLoggedIn"] = false
		if HasUser(c) {
			viewContext["isLoggedIn"] = true
			viewContext["userName"] = GetUser(c).Name
		}
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
