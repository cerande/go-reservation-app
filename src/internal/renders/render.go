package renders

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/cerande/go-reservation-app/src/internal/config"
	"github.com/cerande/go-reservation-app/src/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewTemplate(config *config.AppConfig) {
	app = config
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(req *http.Request, wr http.ResponseWriter, tmplName string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = BuildCache()
	}

	t, ok := tc[tmplName]
	if !ok {
		log.Fatal("Template not found in cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, req)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(wr)
	if err != nil {
		log.Fatal(err)
	}

}

func BuildCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template, 0)

	templateFiles, err := filepath.Glob("./src/templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, file := range templateFiles {
		name := filepath.Base(file)

		ts, err := template.New(name).ParseFiles(file)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./src/templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			for _, m := range matches {
				ts, err = ts.ParseGlob(m)
				if err != nil {
					log.Fatal(err)
				}
				cache[name] = ts
			}
		}
	}
	return cache, nil
}
