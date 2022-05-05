package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

<<<<<<< HEAD
	// panic("Not yet implemented") // TODO: answer here
	// c := &Pokemon{
	// 	Name: "bulbasaur",
	// 	Species: struct {
	// 		Name string "json:\"name\""
	// 	}{
	// 		Name: "bulbasaur",
	// 	},
	// 	Abilities: []struct {
	// 		Ability struct {
	// 			Name string "json:\"name\""
	// 		}
	// 	}{
	// 		struct {
	// 			Ability struct {
	// 				Name string "json:\"name\""
	// 			}
	// 		}{
	// 			struct {
	// 				Name string "json:\"name\""
	// 			}{
	// 				Name: "overgrow",
	// 			},
	// 		},
	// 		struct {
	// 			Ability struct {
	// 				Name string "json:\"name\""
	// 			}
	// 		}{
	// 			struct {
	// 				Name string "json:\"name\""
	// 			}{
	// 				Name: "chlorophyll",
	// 			},
	// 		},
	// 	},
	// }

	// return c, nil
	client := &http.Client{}
	res, err := client.Get(apiPath)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	pokemon := Pokemon{}

	json.NewDecoder(res.Body).Decode(&pokemon)

	return &pokemon, nil
=======
	panic("Not yet implemented") // TODO: answer here
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func main() {
	pokemon, err := GetPokemonData()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pokemon)
}
