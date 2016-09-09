package models

import (
	"mgws/pokedex/pagination"
)

type PType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Types struct {
	Types                 []PType `json:"data"`
	pagination.Pagination `json:"pagination"`
}

type PokedexTypes struct {
	Ids []PType
}

func AllPokeTypes(offset int, limitNo int) (Types, int) {
	rows, _ := db.Query("SELECT * FROM types LIMIT ?,?", offset, limitNo)
	Response := Types{}
	perPage := 0
	for rows.Next() {
		pokeType := PType{}
		rows.Scan(&pokeType.Id, &pokeType.Name)
		Response.Types = append(Response.Types, pokeType)
		perPage++
	}
	if perPage == 0 {
		emptySlice := make([]PType, 0)
		Response.Types = emptySlice
	}
	return Response, perPage

}

func TotalPokemonTypes() int {
	total := 0
	_ = db.QueryRow(`select count(*) from types`).Scan(&total)
	return total
}
