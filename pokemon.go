package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)


func PokemonName(id int, lang string) string{
	file, e := ioutil.ReadFile("./data/" + lang + ".json")
	if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    var pokemons []string
    json.Unmarshal(file, &pokemons)

	return pokemons[id - 1]
}

func PokemonId(name string, lang string) int{
	file, e := ioutil.ReadFile("./data/" + lang + ".json")
	if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    var pokemons []string
    json.Unmarshal(file, &pokemons)

    var pokemonId int

    for i := 0; i < len(pokemons); i++ {
        if pokemons[i] == name {
            pokemonId = i + 1
        }
 	}

	return pokemonId
}