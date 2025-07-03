package repository

import (
	"database/sql"
	"fmt"
	"user-service/models"
)

type UserRepo struct {
	DB *sql.DB
}

func (ur *UserRepo) InitUserTable() error {

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);
	`

	_, err := ur.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	return nil
}

func (ur *UserRepo) CreateUser(user models.User) (int64, error) {

	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	res, err := ur.DB.Exec(query, user.Email, user.Password)

	if err != nil {
		return 0, fmt.Errorf("error while inserting values on DB %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %v", err)
	}

	return id, nil
}

func (ur *UserRepo) GetUser(email string) (models.User, error) {

	var user models.User
	query := "SELECT id, email, password FROM users WHERE email = ?"

	row := ur.DB.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return user, fmt.Errorf("user not found")
	} else if err != nil {
		return user, fmt.Errorf("query error: %v", err)
	}

	return user, nil
}
