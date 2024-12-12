package dto

import (
	"errors"

	"github.com/JoseFredes/learn-go-hex/internal/core/domain"
)

const (
	ErrInvalidEmail    = "invalid email"
	ErrInvalidName     = "invalid name"
	ErrInvalidLastName = "invalid last name"
)

type UserDTO struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

func (u *UserDTO) ValitateNewUser() error {
	if u.Email == "" {
		return errors.New(ErrInvalidEmail)
	}

	if u.Name == "" {
		return errors.New(ErrInvalidName)
	}

	if u.LastName == "" {
		return errors.New(ErrInvalidLastName)
	}

	return nil
}

func (u *UserDTO) ParseToDomain() *domain.User {
	return &domain.User{
		Name:     u.Name,
		LastName: u.LastName,
		Email:    u.Email,
	}
}