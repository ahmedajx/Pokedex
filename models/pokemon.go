package models

import (
	"mgws/pokedex/pagination"
	"strings"
)

type Pokemon struct {
	PokedexID int    `json:"pokedex_id"`
	Name      string `json:"name"`
	types     string
	Types     []string `json:"types"`
}

type CollectionPokemon struct {
	CollectionPokemon     []Pokemon `json:"data"`
	pagination.Pagination `json:"pagination"`
}

func AllPokemons(offset int, limitNo int) (CollectionPokemon, int) {
	rows, _ := db.Query(
		`SELECT  pokemon.*  , group_concat( DISTINCT types.name) as types FROM pokedex.pokemon
		LEFT JOIN pokemon_type
		ON pokemon.pokedexID =  pokemon_type.pokemonID
		LEFT JOIN pokedex.types
		ON pokemon_type.type_id =  types.id
		group by pokemon.name, pokemon.pokedexID
		order by pokedexID
		LIMIT ?,?`, offset, limitNo)
	Response := CollectionPokemon{}
	perPage := 0
	for rows.Next() {
		pokemon := Pokemon{}
		rows.Scan(&pokemon.PokedexID, &pokemon.Name, &pokemon.types)
		types := strings.Split(pokemon.types, ",")
		if pokemon.types == "" {
			types = make([]string, 0)
		}
		pokemon.Types = types
		Response.CollectionPokemon = append(Response.CollectionPokemon, pokemon)
		perPage++
	}
	if perPage == 0 {
		emptySlice := make([]Pokemon, 0)
		Response.CollectionPokemon = emptySlice
	}
	return Response, perPage
}

func TotalPokemons() int {
	total := 0
	_ = db.QueryRow(`select count(*) from pokemon`).Scan(&total)
	return total
}
