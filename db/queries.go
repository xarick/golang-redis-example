package db

import "github.com/xarick/golang-redis-example/models"

func RegisterUser(username, hashedPassword string) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := DB.Exec(query, username, hashedPassword)
	return err
}

func GetUserByName(username string) (*models.User, error) {
	query := "SELECT username, password FROM users WHERE username = $1"
	var user models.User
	err := DB.Get(&user, query, username)
	return &user, err
}
