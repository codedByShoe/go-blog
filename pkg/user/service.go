package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user *User) error
	Login(username string, password string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Register(user *User) error {
	existingUser, err := s.repo.GetUserByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already taken")
	}
	return s.repo.CreateUser(user)
}

func (s *service) Login(username string, password string) (*User, error) {
	user, err := s.repo.GetUserByUsername(username)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Incorrect Password")
	}
	return user, nil
}
