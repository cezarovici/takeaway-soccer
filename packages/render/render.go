package render

import (
	"fmt"
	"html/template"
	"net/http"
)

const templatePath = "./pages/"

func RenderTemplate(w http.ResponseWriter, templateName string) error {
	parsedTemplate, errorParsing := template.ParseFiles(templatePath + templateName)
	if errorParsing != nil {
		fmt.Print("error from parsing the files")
		return errorParsing
	}

	return parsedTemplate.Execute(w, nil)
}

// func RenderTemplateTest(w http.ResponseWriter, tmpl string) error {
// 	myCache := map[string]*template.Template

// 	pages, err := filepath.Glob("./pages/*htm")
// 	if err != nil {
// 		return err
// 	}

// }
