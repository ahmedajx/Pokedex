package main

import (
	"Pokedex/models"
	"Pokedex/routes"
	"net/http"
)

func main() {
	models.Connect()
	router := routes.Router()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":3000", router)
}
