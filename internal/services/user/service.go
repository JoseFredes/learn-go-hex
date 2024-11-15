package services

import (
	"github.com/JoseFredes/learn-go-hex/internal/core/domain"
	"github.com/JoseFredes/learn-go-hex/internal/core/ports"
)

type userServices struct {
	repository ports.UserRepository
}

func NewUserServices(repo ports.UserRepository) *userServices {
	return &userServices{repository: repo}
}

func (s *userServices) Get(id int) (domain.User , error){
	return s.repository.GetUserByID(id)
}