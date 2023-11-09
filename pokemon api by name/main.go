package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	pikachu    = "pikachu"
	charmander = "charmander"
	bulbasaur  = "bulbasaur"
	squirtle   = "squirtle"
	pidgey     = "pidgey"
	rattata    = "rattata"

	expectedPikachuHeight = 4
	expectedPikachuWeight = 60
)

type pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

// if the name matches the json then this doesn't need to be defined. ie name instead o
// f pokemon_name or something.

func getPokemonByName(name string) (pokemon, error) {
	const baseURL = "https://pokeapi.co/api/v2/pokemon/"
	singlePokemon := pokemon{}

	url := baseURL + name

	res, err := http.Get(url)
	if err != nil {
		return singlePokemon, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return singlePokemon, errors.New(fmt.Sprintf("Pokemon %s not found", name))
	}
	if res.StatusCode != http.StatusOK {
		return singlePokemon, errors.New(fmt.Sprintf("unexpected status code %d!", res.StatusCode))
	}

	err = json.NewDecoder(res.Body).Decode(&singlePokemon)
	if err != nil {
		return singlePokemon, err
	}

	return singlePokemon, nil
}

type pokemonMap map[string]interface{}

func getPokemonMapByName(name string) (pokemonMap, error) {
	const baseURL = "https://pokeapi.co/api/v2/pokemon/"
	singlePokemon := pokemonMap{}

	url := baseURL + name

	res, err := http.Get(url)
	if err != nil {
		return singlePokemon, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return singlePokemon, errors.New(fmt.Sprintf("Pokemon %s not found", name))
	}
	if res.StatusCode != http.StatusOK {
		return singlePokemon, errors.New(fmt.Sprintf("unexpected status code %d!", res.StatusCode))
	}

	err = json.NewDecoder(res.Body).Decode(&singlePokemon)
	if err != nil {
		return singlePokemon, err
	}

	return singlePokemon, nil
}

func pokemonPrompt() string {
	fmt.Println("please enter the name of a pokemon")
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {

	// TODO: extract get logic into a function that takes a pokemon name and returns a pokemon stuct

	pokemonName := pokemonPrompt()

	pokemonData, err := getPokemonByName(pokemonName)
	//pokemonData, err := getPokemonMapByName(pokemonName)
	if err != nil {
		log.Fatal(err)
	}

	//unmarshal turns a string into a struct
	// marshal changes a struct into a string

	// another way to do this is with newDecoder ie
	// Decode the JSON response
	//	var pikachuData []pokemon
	//	err = json.NewDecoder(response.Body).Decode(&pikachuData)
	//	if err != nil {
	//		return nil, err
	//	}

	// for k, _ := range pokemonData {
	// 	fmt.Printf("%s\n", k)
	// }

	fmt.Println(pokemonData)

	fmt.Printf("%s's height: %d\n", pokemonName, pokemonData.Height)
	fmt.Printf("%s's weight: %d\n", pokemonName, pokemonData.Weight)

	helloWorld := "Hello World!"

	// sl []struct{} // 128MB

	fmt.Println(f(helloWorld))
	fmt.Println(helloWorld)
}

func f(in string) string {
	in = "Goodbye World!"
	return in
}

/*
// https://pokeapi.co/api/v2/pokemon/pikachu

Part 1
Write a script that will make a request to the PokeAPI (https://pokeapi.co) to get the following data points for “pikachu” and then print them to the terminal:

height
weight


Part 2
Update your script to pull the same data points for “pikachu”, “charmander”, “bulbasaur”, “squirtle”, “pidgey”, and “rattata”, and output the following summary statistics:

mean height
median height
mode height
mean weight
median weight
mode weight
The summary statistics should be calculated by hand (i.e. not using a package/library).

For reference:

Mean = the average of a list of values

[1, 2, 3] -> 2
[1, 3, 4] -> 2.66667
Median = the middle value of a list of sorted values (if there is an even number of values, the average of the 2 middle numbers is considered the median)

[1, 2, 3] -> 2
[1, 2, 3, 4] -> 2.5
Mode = the most common value(s) of a list of values

[1, 1, 2, 2, 2] -> [2]
[1, 1, 1, 2, 2, 2] -> [1, 2]
[1, 2, 3, 4, 5] -> [1, 2, 3, 4, 5]
Some notes about the Go environment in Coderpad:

Only the go standard library is available
Only running main is avaible. go test, etc. is not available
All output should be printed to the console
*/
