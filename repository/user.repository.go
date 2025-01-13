package repository

import (
	"database/sql"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/db"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id"
	err := db.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, email, password, created_at FROM users WHERE username = $1"
	row := db.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
