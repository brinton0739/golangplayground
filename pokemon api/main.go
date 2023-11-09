package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func main() {

	// TODO: extract get logic into a function that takes a pokemon name and returns a pokemon stuct

	const url = "https://pokeapi.co/api/v2/pokemon/pikachu"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// defer will close the data after main has finished running.

	pokeData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	pikachuData := pokemon{}
	err = json.Unmarshal(pokeData, &pikachuData)
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

	fmt.Println(pikachu)

	fmt.Printf("Pikachu's height: %d\n", pikachuData.Height)
	fmt.Printf("Pikachu's weight: %d\n", pikachuData.Weight)
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
