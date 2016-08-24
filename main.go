package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mgws/pokedex/models"
	"net/http"
	//"strconv"
)

func pokedexIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pragma", "no-cache")
	response, perPage := models.AllPokemons()
	response.Pagination.Total = models.TotalPokemons()
	response.Pagination.PerPage = perPage
	b, _ := json.Marshal(response)
	w.Write(b)
}

func main() {
	models.Connect()
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/pokedex", pokedexIndex).Methods("GET")
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3067", nil)
}
