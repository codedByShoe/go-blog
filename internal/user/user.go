package user

import (
	"gorm.io/gorm"
)

type User struct {
	ID       int    `db:"id" gorm:"primaryKey"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type Model interface {
	// Define Methods
	CreateUser(user *User) error
	GetUserByUsername(username string) (User, error)
}

type userModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) Model {
	return &userModel{db: db}
}

func (m *userModel) CreateUser(user *User) error {
	// TODO: hash password before passing to db
	tx := m.db.Create(user)

	return tx.Error
}

func (m *userModel) GetUserByUsername(username string) (User, error) {
	var user User
	tx := m.db.Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}
	return user, nil
}
