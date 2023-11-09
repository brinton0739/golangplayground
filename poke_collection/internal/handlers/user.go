package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"poke_collection/internal/database"
	"strconv"

	"github.com/gorilla/mux"
)

/*
	GET -- Get a single or list
	POST -- Create
	DELETE -- Remove
	PUT -- Change overwrite everything with what I have
	PATCH -- Cange only what I send
*/

func (p PokeHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := database.User{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "unable to parse json", http.StatusInternalServerError)
		fmt.Println("unable to marshal json", err.Error())
		return
	}

	newUser, err := p.db.CreateUser(data)
	if err != nil {
		http.Error(w, "unable to create user", http.StatusInternalServerError)
		fmt.Println("unable to create user", err.Error())
		return
	}
	output, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		fmt.Println("unable to marshal json", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(output))
}

func (p PokeHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	realID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	user, err := p.db.GetUserByID(realID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error getting user")
		return
	}
	j, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error marshalling user")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(j))
}

func (p PokeHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	realID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	var updateUser database.User

	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error parsing request body")
		return
	}

	if updateUser.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is required")
		return
	}

	updateUser.ID = realID

	updateUser, err = p.db.UpdateUserByID(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error updating user")
		return
	}

	defer r.Body.Close()

	j, err := json.Marshal(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error marshalling user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(j))
}

func (p PokeHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	realID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID")
		return
	}

	err = p.db.DeleteUserByID(realID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error deleting user")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// curl -X PUT http://localhost:8080/users/3 -H "Content-Type: application/json" -d '{"id": 3, "name":"Brinton", "password":"password1", "email":"brinton@updated" }'
