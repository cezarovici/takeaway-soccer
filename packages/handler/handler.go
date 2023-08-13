package handler

import (
	"encoding/json"
	"log"
	"net/http"

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

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) HandleAdaugaEditieJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Hello Cezar",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func (m *Repository) PostHandleAddEditie(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	informatiiEditie := model.Editie{
		Data:  r.Form.Get("data_editie"),
		Numar: r.Form.Get("numar_editie"),
	}

	form := form.New(r.PostForm)

	form.Has("data_editie", r)
	form.Has("numar_editie", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["informatiiEditie"] = informatiiEditie

		errRendingTemplate := render.RenderTemplate(w, r, "adauga_etapa_page.htm", &model.TemplateData{
			Form: form,
			Data: data,
		})
		if errRendingTemplate != nil {
			log.Print(errRendingTemplate)
		}
	}

}
