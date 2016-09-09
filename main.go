package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mgws/pokedex/models"
	"mgws/pokedex/pagination"
	"net/http"
	"strconv"
)

func pokeTypesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pragma", "no-cache")
	offset, limitNo, pageNo := pagination.Paginate(r)
	response, perPage := models.AllPokeTypes(offset, limitNo)
	response.Pagination.Total = models.TotalPokemonTypes()
	response.Pagination.PerPage = perPage
	response.Pagination.PageNo = pageNo
	b, _ := json.Marshal(response)
	w.Write(b)
}

func pokedexCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newPokemon := models.Pokemon{}
	var pokemon models.Pokemon
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &pokemon)
	newPokemon.PokedexID = pokemon.PokedexID
	newPokemon.Name = pokemon.Name
	models.CreatePokemon(newPokemon)
	output, _ := json.Marshal(newPokemon)
	w.Write(output)
}

func pokedexIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pragma", "no-cache")
	offset, limitNo, pageNo := pagination.Paginate(r)
	include := r.FormValue("include")
	response, perPage := models.AllPokemons(offset, limitNo, include)
	response.Pagination.Total = models.TotalPokemons()
	response.Pagination.PerPage = perPage
	response.Pagination.PageNo = pageNo
	b, _ := json.Marshal(response)
	w.Write(b)
}

func pokedexPokeTypeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := mux.Vars(r)
	pokemonId := urlParams["pokemonID"]
	pokemon := models.Pokemon{}
	i, _ := strconv.Atoi(pokemonId)
	pokemon.PokedexID = i
	z := models.IncludePokeTypes(pokemon)
	b, _ := json.Marshal(z)
	w.Write(b)
}

func main() {
	models.Connect()
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/poke_types", pokeTypesIndex).Methods("GET")
	gorillaRoute.HandleFunc("/api/pokedex", pokedexIndex).Methods("GET")
	gorillaRoute.HandleFunc("/api/pokedex", pokedexCreate).Methods("POST")
	gorillaRoute.HandleFunc("/api/pokedex/{pokemonID:[0-9]+}/poke_types", pokedexPokeTypeIndex).Methods("GET")
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3000", nil)
}
