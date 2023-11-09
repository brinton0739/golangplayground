package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poke_collection/internal/database"
	"strconv"

	"github.com/gorilla/mux"
)

// users/{id} -- > GETS UserByID
// users/{user_id}/pokemon/{pokemon_id} --> GETS User and Pokemon by UserID
// users/{user_id}/pokemon --> POSTS Pokemon to User {user_id: 1, pokemon_name: "Pikachu"}

func (p PokeHandler) InsertPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars["user_id"]

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	data := database.Pokemon{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "unable to parse json", http.StatusInternalServerError)
		fmt.Println("unable to marshal json", err.Error())
		return
	}

	data.UserID = userID

	newPokemon, err := p.db.InsertPokemon(data)
	if err != nil {
		http.Error(w, "unable to insert pokemon", http.StatusInternalServerError)
		fmt.Println("unable to insert pokemon", err.Error())
		return
	}
	output, err := json.Marshal(newPokemon)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		fmt.Println("unable to marshal json", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(output))

}

func (p PokeHandler) GetPokemonByUserID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userIDString := vars["user_id"]

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	pokemon, err := p.db.GetPokemonByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error getting user")
		return
	}

	j, err := json.Marshal(pokemon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error marshalling user")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(j))

}

func (p PokeHandler) UpdatePokemonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	realID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	var updatePokemon database.Pokemon

	err = json.NewDecoder(r.Body).Decode(&updatePokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error parsing request body")
		return
	}

	updatePokemon.ID = realID

	updatePokemon, err = p.db.UpdatePokemonByID(updatePokemon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error updating pokemon")
		return
	}

	defer r.Body.Close()

	j, err := json.Marshal(updatePokemon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error marshaling pokemon")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(j))

}

func (p PokeHandler) DeletePokemonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	realID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	err = p.db.DeletePokemonByID(realID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error deleting user")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
