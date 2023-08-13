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

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Has(field string, r *http.Request) bool {
	content := r.Form.Get(field)
	if content == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}
