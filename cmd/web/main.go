package main

import (
	"net/http"

	handler "github.com/cezarovici/takeaway-soccer/packages/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handler.HandeHomePage)
	http.HandleFunc("/jucatori", handler.HandePlayersPage)
	http.HandleFunc("/etape", handler.HandeEtapePage)

	http.ListenAndServe(portNumber, nil)
}
