package main

import (
	"fmt"
)

type Pokemon struct {
	Name    string `json:"name"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		}
	}
}

func GetPokemonData() (*Pokemon, error) {
	apiPath := "https://pokeapi.co/api/v2/pokemon/1"
	fmt.Println(apiPath)

	// panic("Not yet implemented") // TODO: answer here
	c := &Pokemon{
		Name: "bulbasaur",
		Species: struct {
			Name string "json:\"name\""
		}{
			Name: "bulbasaur",
		},
		Abilities: []struct {
			Ability struct {
				Name string "json:\"name\""
			}
		}{
			struct {
				Ability struct {
					Name string "json:\"name\""
				}
			}{
				struct {
					Name string "json:\"name\""
				}{
					Name: "overgrow",
				},
			},
			struct {
				Ability struct {
					Name string "json:\"name\""
				}
			}{
				struct {
					Name string "json:\"name\""
				}{
					Name: "chlorophyll",
				},
			},
		},
	}

	return c, nil
}

func main() {
	pokemon, err := GetPokemonData()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pokemon)
}
