package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	// Define Methods
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// insert the user into the database
	_, err = r.db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	return err
}

func (r *userRepository) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT id, username, email FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
