package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)


func getName(id int, lang string) string{
	file, e := ioutil.ReadFile("./data/" + lang + ".json")
	if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    var pokemons []string
    json.Unmarshal(file, &pokemons)

	return pokemons[id - 1]
}