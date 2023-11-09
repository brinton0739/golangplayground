package handlers

import (
	"fmt"
	"net/http"
	"poke_collection/internal/database"
)

type PokeHandler struct {
	db database.PokeDB
}

func NewPokeHandler(dbIn database.PokeDB) *PokeHandler {
	return &PokeHandler{db: dbIn}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Darkness")
	fmt.Fprint(w, "Hello Darkness")
}
