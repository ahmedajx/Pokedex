package handlers

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"mgws/pokedex/errors"
	"mgws/pokedex/models"
	"mgws/pokedex/pagination"
	"net/http"
	"strconv"
)

func PokeTypesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	offset, limitNo, pageNo := pagination.Paginate(r)
	response, perPage := models.AllPokeTypes(offset, limitNo)
	response.Pagination.Total = models.TotalPokemonTypes()
	response.Pagination.PerPage = perPage
	response.Pagination.PageNo = pageNo
	b, _ := json.Marshal(response)
	w.Write(b)
}
func Index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.gohtml")
	err := tpl.Execute(w, "OH")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
func GetPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := mux.Vars(r)
	id, _ := strconv.Atoi(urlParams["pokemonID"])
	err, getPokemon := models.GetSinglePokemon(id)
	switch {
	case err == sql.ErrNoRows:
		errorNotFound := errors.Error{"Not Found", 404}
		w.WriteHeader(http.StatusNotFound)
		output, _ := json.Marshal(errorNotFound)
		w.Write(output)
	case err != nil:
		internalServerError := errors.Error{"Error", 500}
		w.WriteHeader(http.StatusInternalServerError)
		output, _ := json.Marshal(internalServerError)
		w.Write(output)
	default:
		output, _ := json.Marshal(getPokemon)
		w.Write(output)
	}
}
func PokedexCreate(w http.ResponseWriter, r *http.Request) {
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

func PokedexIndex(w http.ResponseWriter, r *http.Request) {
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

func PokedexPokeTypeCreate(w http.ResponseWriter, r *http.Request) {
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

func PokedexPokeTypeIndex(w http.ResponseWriter, r *http.Request) {
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
