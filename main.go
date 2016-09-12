package main

import (
	"mgws/pokedex/models"
	"mgws/pokedex/routes"
	"net/http"
)

func main() {
	models.Connect()
	router := routes.ApiRouter()
	http.ListenAndServe(":3000", router)
}
