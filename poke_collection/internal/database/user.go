package database

type User struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Email    string `db:"email" json:"email"`
}

func (p PokeDB) CreateUser(user User) (User, error) {
	var newID int64
	err := p.db.QueryRow("INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Password, user.Email).
		Scan(&newID)
	if err != nil {
		return User{}, err
	}
	user.ID = newID
	return user, nil
}

func (p PokeDB) GetUserByID(id int64) (User, error) {
	var user User
	err := p.db.QueryRow("SELECT id, name, password, email FROM users where id = $1", id).
		Scan(&user.ID, &user.Name, &user.Password, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (p PokeDB) UpdateUserByID(user User) (User, error) {

	var newUser User
	err := p.db.QueryRow("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4 RETURNING id, name, password, email;", user.Name, user.Email, user.Password, user.ID).
		Scan(&newUser.ID, &newUser.Name, &newUser.Password, &newUser.Email)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (p PokeDB) DeleteUserByID(id int64) error {
	_, err := p.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil

}
