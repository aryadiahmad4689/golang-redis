package main

import (
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/pokemonwithredis", getPokemonWithRedis)
	// http.HandleFunc("/pokemonwithoutredis", getPokemonWithoutRedis)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
