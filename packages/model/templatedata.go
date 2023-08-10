package model

import "github.com/cezarovici/takeaway-soccer/packages/form"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}

	CSRFToken string
	Flush     string
	Warning   string
	Error     string

	Form *form.Form
}
