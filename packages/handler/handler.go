package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/araddon/dateparse"
	"github.com/cezarovici/takeaway-soccer/packages/config"
	"github.com/cezarovici/takeaway-soccer/packages/form"
	"github.com/cezarovici/takeaway-soccer/packages/model"
	"github.com/cezarovici/takeaway-soccer/packages/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HandleHomePage(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Dev!"

	errRendingTemplate := render.RenderTemplate(w, r, "home_page.htm", &model.TemplateData{
		StringMap: stringMap,
	})
	if errRendingTemplate != nil {
		log.Print(errRendingTemplate)
	}
}

func (m *Repository) HandleAddEditie(w http.ResponseWriter, r *http.Request) {
	errRendingTemplate := render.RenderTemplate(w, r, "adauga_etapa_page.htm", &model.TemplateData{
		Form: form.New(nil),
	})
	if errRendingTemplate != nil {
		log.Print(errRendingTemplate)
	}
}

func (m *Repository) PostHandleAddEditie(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	editie_nr, err := strconv.Atoi(r.Form.Get("numar_editie"))
	data_nr, err := dateparse.ParseAny(r.Form.Get("data_editie"))

	log.Print(editie_nr, data_nr)
}
