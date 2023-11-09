package main

import (
	"fmt"
	"log"
	"net/http"
	"poke_collection/internal/database"
	"poke_collection/internal/handlers"

	"github.com/gorilla/mux"
)

// server > handlers > database

func main() {
	db, err := database.NewPokeDB()
	if err != nil {
		log.Fatal("Error creating database")
	}

	defer db.Close()

	pokeHandler := handlers.NewPokeHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HelloWorld)
	router.HandleFunc("/users/{id}", pokeHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users", pokeHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", pokeHandler.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/users/{id}", pokeHandler.DeleteUserByID).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/pokemon", pokeHandler.GetPokemonByUserID).Methods("GET")
	router.HandleFunc("/users/{user_id}/pokemon", pokeHandler.InsertPokemon).Methods("POST")
	router.HandleFunc("/users/{user_id)/pokemon/{id}", pokeHandler.UpdatePokemonByID).Methods("PUT")
	router.HandleFunc("/users/{user_id}/pokemon/{id}", pokeHandler.DeletePokemonByID).Methods("DELETE")
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", router)

}
