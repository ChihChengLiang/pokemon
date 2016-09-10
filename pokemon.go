package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func PokemonName(id int, lang string) string {
	path, _ := filepath.Abs("./data/" + lang + ".json")
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var pokemons []string
	json.Unmarshal(file, &pokemons)

	return pokemons[id-1]
}

func PokemonId(name string, lang string) int {
	path, _ := filepath.Abs("./data/" + lang + ".json")
	file, e := ioutil.ReadFile(path)
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
