package routes

import (
	"mgws/pokedex/auth"
	"mgws/pokedex/handlers"
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
	Route{"GET", "/pokedex", handlers.PokedexIndex},
	Route{"POST", "/pokedex", handlers.PokedexCreate},
	Route{"GET", "/poke_types", handlers.PokeTypesIndex},
	Route{"GET", "/pokedex/{pokemonID:[0-9]+}/poke_types", handlers.PokedexPokeTypeIndex},
	Route{"POST", "/pokedex/{pokemonID:[0-9]+}/poke_types", auth.Middleware(handlers.PokedexPokeTypeCreate)},
	Route{"POST", "/pokedex", handlers.PokedexCreate},
	Route{"GET", "/auth", auth.Auth},
}
var webRoutes = WebRoutes{
	Route{"GET", "/", handlers.Index},
}
