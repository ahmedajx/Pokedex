package routes

import (
	"Pokedex/auth"
	"Pokedex/handlers"
	"net/http"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Apiroutes []Route

type WebRoutes []Route

var apiRoutes = Apiroutes{
	Route{"GET", "/auth", auth.Auth},
	Route{"GET", "/pokedex", handlers.PokedexIndex},
	Route{"GET", "/poke_types", handlers.PokeTypesIndex},
	Route{"GET", "/pokedex/{pokemonID:[0-9]+}/poke_types", handlers.PokedexPokeTypeIndex},
	Route{"GET", "/pokedex/{pokemonID:[0-9]+}", handlers.GetPokemon},
	Route{"POST", "/pokedex", auth.Middleware(handlers.PokedexCreate)},
	Route{"POST", "/pokedex/{pokemonID:[0-9]+}/poke_types", auth.Middleware(handlers.PokedexPokeTypeCreate)},
}
var webRoutes = WebRoutes{
	Route{"GET", "/", handlers.Index},
}
