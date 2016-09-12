package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mgws/pokedex/auth"
	"mgws/pokedex/models"
	"mgws/pokedex/pagination"
	"net/http"
	"strconv"
)

func pokeTypesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

func pokedexPokeTypeCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := ioutil.ReadAll(r.Body)
	ids := make([]models.PType, 0)
	json.Unmarshal(b, &ids)
	urlParams := mux.Vars(r)
	pokemonId, _ := strconv.Atoi(urlParams["pokemonID"])
	for _, v := range ids {
		models.SavePokemonType(pokemonId, v.Id)
	}
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
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/poke_types", pokeTypesIndex).Methods("GET")
	subRouter.HandleFunc("/pokedex", pokedexIndex).Methods("GET")
	subRouter.HandleFunc("/pokedex", auth.Middleware(pokedexCreate)).Methods("POST")
	subRouter.HandleFunc("/pokedex/{pokemonID:[0-9]+}/poke_types", pokedexPokeTypeIndex).Methods("GET")
	subRouter.HandleFunc("/pokedex/{pokemonID:[0-9]+}/poke_types", auth.Middleware(pokedexPokeTypeCreate)).Methods("POST")
	subRouter.HandleFunc("/auth", auth.Auth).Methods("GET")
	http.ListenAndServe(":3000", router)
}
