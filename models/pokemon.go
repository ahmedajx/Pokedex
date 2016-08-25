package models

type Pokemon struct {
	PokedexID int    `json:"pokedex_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type Pagination struct {
	Total   int `json:"total"`
	PerPage int `json:"per_page"`
	PageNo  int `json:"page_no"`
}

type CollectionPokemon struct {
	CollectionPokemon []Pokemon `json:"data"`
	Pagination        `json:"pagination"`
}

func AllPokemons(offset int, limitNo int) (CollectionPokemon, int) {
	rows, _ := db.Query("SELECT * FROM pokemon LIMIT ?,?", offset, limitNo)
	Response := CollectionPokemon{}
	perPage := 0
	for rows.Next() {
		pokemon := Pokemon{}
		rows.Scan(&pokemon.PokedexID, &pokemon.Name, &pokemon.Type)
		Response.CollectionPokemon = append(Response.CollectionPokemon, pokemon)
		perPage++
	}
	return Response, perPage
}

func TotalPokemons() int {
	total := 0
	_ = db.QueryRow(`select count(*) from pokemon`).Scan(&total)
	return total
}
