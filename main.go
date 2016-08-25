package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mgws/pokedex/models"
	"mgws/pokedex/pagination"
	"net/http"
)

func pokedexIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pragma", "no-cache")
	offset, limitNo, pageNo := pagination.Paginate(r)
	response, perPage := models.AllPokemons(offset, limitNo)
	response.Pagination.Total = models.TotalPokemons()
	response.Pagination.PerPage = perPage
	response.Pagination.PageNo = pageNo
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
