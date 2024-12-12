package services

import (
	"github.com/JoseFredes/learn-go-hex/internal/core/domain"
	"github.com/JoseFredes/learn-go-hex/internal/core/ports"
)

type UserServices interface {
	Get(id int) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

type userServices struct {
	repository ports.UserRepository
}

func NewUserServices(repository ports.UserRepository) UserServices {
	return &userServices{repository: repository}
}

func (s *userServices) Get(id int) (*domain.User, error) {
	return s.repository.GetUserByID(id)
}

func (s *userServices) Create(user *domain.User) (*domain.User, error) {
	if err := user.ValitateNewUser(); err != nil {
		return nil, err
	}

	return s.repository.CreateUser(user)
}