package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/chrslex/bookings-mini-project/pkg/config"
	"github.com/chrslex/bookings-mini-project/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(cfg *config.AppConfig) {
	app = cfg
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Get the template cache from the app config

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing templates to browser")
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// Search file path with matching pattern
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// Last element of file path
		name := filepath.Base(page)

		// Creating new template with corresponding name, fucntions, and path
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Search for layout files
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// Parses the template definition from the matching layout files
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		// Assign template based on name
		myCache[name] = ts

	}

	return myCache, nil
}
