package models

import (
	"log"
	"mgws/pokedex/pagination"
)

type Pokemon struct {
	PokedexID int    `json:"pokedex_id"`
	Name      string `json:"name"`
	//http://stackoverflow.com/questions/18088294/how-to-not-marshal-an-empty-struct-into-json-with-go
	Relationship *Relationship `json:"relationship,omitempty"`
}

type Relationship struct {
	Types []PType `json:"types"`
}

type CollectionPokemon struct {
	CollectionPokemon     []Pokemon `json:"data"`
	pagination.Pagination `json:"pagination"`
}

func AllPokemons(offset int, limitNo int, include string) (CollectionPokemon, int) {
	rows, _ := db.Query("SELECT * from pokemon  order by pokedexID LIMIT ?,?", offset, limitNo)
	Response := CollectionPokemon{}
	perPage := 0
	for rows.Next() {
		pokemon := Pokemon{}
		rows.Scan(&pokemon.PokedexID, &pokemon.Name)
		if include == "types" {
			pokemon.Relationship = IncludePokeTypes(pokemon)
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

func CreatePokemon(newPokemon Pokemon) {
	stmt, err := db.Prepare("INSERT INTO pokemon(pokedexID,name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(newPokemon.PokedexID, newPokemon.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func IncludePokeTypes(pokemon Pokemon) *Relationship {
	r := &Relationship{}
	c := 0
	x := true
	pokeTyperows, _ := db.Query(`SELECT 
		    distinct types.id, types.name
		FROM
		    pokedex.pokemon
		        LEFT JOIN
		    pokemon_type ON pokemon.pokedexID = pokemon_type.pokemonID
		        LEFT JOIN
		    pokedex.types ON pokemon_type.type_id = types.id
		where pokedexID = ?`, pokemon.PokedexID)
	for pokeTyperows.Next() {
		c++
		ptypes := PType{}
		pokeTyperows.Scan(&ptypes.Id, &ptypes.Name)
		r.Types = append(r.Types, ptypes)
		if ptypes.Name == "" {
			x = false
		}
	}
	if c == 0 || x == false {
		emptySlice := make([]PType, 0)
		r.Types = emptySlice
	}
	return r
}

func TotalPokemons() int {
	total := 0
	_ = db.QueryRow(`select count(*) from pokemon`).Scan(&total)
	return total
}
