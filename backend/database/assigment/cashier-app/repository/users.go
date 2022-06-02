package repository

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {

	var user User
	var sqlStatement string

	sqlStatement = "SELECT * FROM users WHERE id = ?"

	row := u.db.QueryRow(sqlStatement, id)

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin); err != nil {
		return user, err
	}

	return user, nil
	// return User{}, nil // TODO: replace this
}

func (u *UserRepository) FetchUsers() ([]User, error) {

	var users []User
	var sqlStatement string

	sqlStatement = "SELECT * FROM users"

	rows, err := u.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
	// return []User{}, nil // TODO: replace this
}

func (u *UserRepository) Login(username string, password string) (*string, error) {

	// var user User
	// var sqlStatement string

	// sqlStatement = "SELECT * FROM users WHERE username = ? AND password = ?"

	// row := u.db.QueryRow(sqlStatement, username, password)

	// if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin); err != nil {
	// 	return nil, err
	// }

	// if user.Loggedin == true {
	// 	return nil, errors.New("User is already logged in")
	// }

	// sqlStatement = "UPDATE users SET loggedin = ? WHERE username = ?"

	// _, err := u.db.Exec(sqlStatement, true, username)

	// if err != nil {
	// 	return nil, err
	// }

	// return &user.Username, nil
	var sqlStatement string

	sqlStatement = "SELECT * FROM users WHERE username = ? AND password = ?"

	row := u.db.QueryRow(sqlStatement, username, password)
	user := User{}

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)

	if err != nil {
		return &user.Username, errors.New("Login Failed")
	}

	if user.Loggedin == true {
		return nil, errors.New("User is already logged in")
	}

	sqlStatement = "UPDATE users SET loggedin = ? WHERE username = ?"

	_, err = u.db.Exec(sqlStatement, true, username)

	if err != nil {
		return nil, err
	}

	return &user.Username, nil // TODO: replace this
}

func (u *UserRepository) InsertUser(username string, password string, role string, loggedin bool) error {

	var sqlStatement string

	sqlStatement = "INSERT INTO users (username, password, role, loggedin) VALUES (?, ?, ?, ?)"

	_, err := u.db.Exec(sqlStatement, username, password, role, loggedin)

	if err != nil {
		return err
	}

	return nil
	// return nil // TODO: replace this
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {

	var user User

	var sqlStatement string

	sqlStatement = "SELECT * FROM users WHERE username = ?"

	row := u.db.QueryRow(sqlStatement, username)

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin); err != nil {
		return nil, err
	}

	return &user.Role, nil
	// return nil, nil // TODO: replace this
}
