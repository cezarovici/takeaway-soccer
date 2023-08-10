package form

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(
			map[string][]string{},
		),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	return r.Form.Get(field) != ""
}