package models

import (
	"mgws/pokedex/pagination"
)

type Pokemon struct {
	PokedexID    int          `json:"pokedex_id"`
	Name         string       `json:"name"`
	Relationship Relationship `json:"relationships"`
}

type Relationship struct {
	Types []PType `json:"types"`
}

type CollectionPokemon struct {
	CollectionPokemon     []Pokemon `json:"data"`
	pagination.Pagination `json:"pagination"`
}

func AllPokemons(offset int, limitNo int) (CollectionPokemon, int) {
	rows, _ := db.Query("SELECT * from pokemon  order by pokedexID LIMIT ?,?", offset, limitNo)
	Response := CollectionPokemon{}
	perPage := 0
	//typesNo := 0
	for rows.Next() {
		pokemon := Pokemon{}
		rows.Scan(&pokemon.PokedexID, &pokemon.Name)
		rozs, _ := db.Query(`SELECT 
		    types.id, types.name
		FROM
		    pokedex.pokemon
		        LEFT JOIN
		    pokemon_type ON pokemon.pokedexID = pokemon_type.pokemonID
		        LEFT JOIN
		    pokedex.types ON pokemon_type.type_id = types.id
		where pokedexID = ?`, pokemon.PokedexID)
		for rozs.Next() {
			ptypes := PType{}
			rozs.Scan(&ptypes.Id, &ptypes.Name)
			pokemon.Relationship.Types = append(pokemon.Relationship.Types, ptypes)
			if ptypes.Name == "" {
				emptySlice := make([]PType, 0)
				pokemon.Relationship.Types = emptySlice
			}
		}
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
