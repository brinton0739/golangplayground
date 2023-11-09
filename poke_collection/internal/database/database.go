package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PokeDB struct {
	db *sql.DB
}

func NewPokeDB() (PokeDB, error) {
	connectionString := "host=localhost port=5432 user=postgres password=postgres dbname=poke_collection sslmode=disable"
	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		return PokeDB{}, err
	}

	// pokeDB := PokeDB{}
	// pokeDB.db = connection
	// return pokeDB, nil

	return PokeDB{db: connection}, nil
}

func (p PokeDB) Close() error {
	return p.db.Close()
}

func (p PokeDB) Ping() error {
	return p.db.Ping()
}
