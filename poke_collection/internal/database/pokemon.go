package database

type Pokemon struct {
	ID     int64  ` db:"id" json:"id"`
	UserID int64  `db:"user_id" json:"user_id"`
	Name   string `db:"pokemon_name" json:"name"`
}

func (p PokeDB) InsertPokemon(pokemon Pokemon) (Pokemon, error) {
	var newPokemon Pokemon

	err := p.db.QueryRow("INSERT INTO users (user_id, pokemon_name) VALUES ($1, $2) RETURNING id, user_id, pokemon_name", pokemon.UserID, pokemon.Name).
		Scan(&newPokemon.ID, &newPokemon.UserID, &newPokemon.Name)
	if err != nil {
		return Pokemon{}, err
	}
	return newPokemon, nil
}

func (p PokeDB) GetPokemonByUserID(id int64) (Pokemon, error) {
	var pokemon Pokemon
	err := p.db.QueryRow("SELECT id, user_id, name FROM pokemon id = $1", id).
		Scan(&pokemon.ID, &pokemon.UserID, &pokemon.Name)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}

func (p PokeDB) UpdatePokemonByID(pokemon Pokemon) (Pokemon, error) {

	var updatePokemon Pokemon
	err := p.db.QueryRow("UPDATE pokemon SET id = $1, user_id = $2, pokemon_name = $3 RETURNING id, user_id, pokemon_name", pokemon.ID, pokemon.UserID, pokemon.Name).
		Scan(&updatePokemon.ID, &updatePokemon.UserID, &updatePokemon.Name)
	if err != nil {
		return Pokemon{}, err
	}
	return updatePokemon, nil
}

func (p PokeDB) DeletePokemonByID(id int64) error {
	_, err := p.db.Exec("DELETE FROM pokemon WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
