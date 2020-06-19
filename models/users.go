package models

import (
	"log"

	"github.com/rodolfobandeira/go-bcrypt-brute/db"
)

// User struct
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
	PasswordSalt      string
}

// GetUsers - Retrieve all non-archived users from the database
func GetUsers() []User {
	db := db.MySQLConnection()

	rows, err := db.Query(
		"SELECT id, email, encrypted_password, password_salt FROM users WHERE is_archived = 0 AND encrypted_password IS NOT NULL AND (role = 'manager' OR role = 'admin')",
	)

	if err != nil {
		log.Fatal(err)
	}

	user := User{}
	users := []User{}

	for rows.Next() {
		var id int
		var email string
		var encryptedPassword string
		var passwordSalt string

		err = rows.Scan(&id, &email, &encryptedPassword, &passwordSalt)
		if err != nil {
			log.Fatal(err)
		}

		user.ID = id
		user.Email = email
		user.EncryptedPassword = encryptedPassword
		user.PasswordSalt = passwordSalt

		users = append(users, user)
	}

	defer db.Close()

	return users
}
