package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user *User) error
	Login(username string, password string) (*User, error)
}

type userService struct {
	repo Repository
}

func NewUserService(r Repository) Service {
	return &userService{repo: r}
}

func (s *userService) Register(user *User) error {
	existingUser, err := s.repo.GetUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already taken")
	}
	return s.repo.CreateUser(user)
}

func (s *userService) Login(username string, password string) (*User, error) {
	user, err := s.repo.GetUserByUsername(username)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Incorrect Password")
	}
	return user, nil
}
