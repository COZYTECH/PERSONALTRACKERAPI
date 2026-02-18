package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/database"
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/models"
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/utils"
)

// RegisterUser registers a new user
func RegisterUser(email, password string) error {
	// Hash the password
	hashed, err := HashPassword(password)
	if err != nil {
		return err
	}

	// Insert user into DB
	_, err = database.DB.Exec(
		"INSERT INTO users (email, password, created_at) VALUES (?, ?, ?)",
		email, hashed, time.Now(),
	)
	if err != nil {
		return errors.New("email may already exist")
	}

	return nil
}

// LoginUser validates credentials and returns a JWT token
func LoginUser(email, password string) (string, error) {
	var user models.User
	err := database.DB.QueryRow(
		"SELECT id, password FROM users WHERE email = ?",
		email,
	).Scan(&user.ID, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	// Check password
	if err := CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
