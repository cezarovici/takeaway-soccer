package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/cezarovici/takeaway-soccer/packages/config"
	"github.com/cezarovici/takeaway-soccer/packages/model"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, td *model.TemplateData) error {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, exists := templateCache[templateName]
	if !exists {
		return fmt.Errorf("the template does not exists")
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buffer, td)

	_, err := buffer.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}

// cretes a template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.htm")
	if err != nil {
		log.Print("the erros if from parsing all pages")
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Print("the erros if from templating all pages")

			return myCache, err
		}

		match, err := filepath.Glob("./templates/*layout.htm")
		if err != nil {
			log.Print("the erros if from matching all pages")

			return myCache, err
		}

		if len(match) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.htm")
			if err != nil {
				log.Print("the erros if from caching all pages")

				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
