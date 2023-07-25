package handler

import (
	"fmt"
	"net/http"

	"github.com/cezarovici/takeaway-soccer/render"
)

func HandeHomePage(w http.ResponseWriter, r *http.Request) {
	if render.RenderTemplate(w, "home.htm") != nil {
		fmt.Print("error rendering the home page")
	}
}

func HandePlayersPage(w http.ResponseWriter, r *http.Request) {
	if render.RenderTemplate(w, "jucatori.htm") != nil {
		fmt.Print("error rendering the jucatori page")
	}
}

func HandeEtapePage(w http.ResponseWriter, r *http.Request) {
	if render.RenderTemplate(w, "etape.htm") != nil {
		fmt.Print("error rendering the etape page")
	}
}
