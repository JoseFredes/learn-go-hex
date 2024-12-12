package ports

import "github.com/JoseFredes/learn-go-hex/internal/core/domain"

// interfaces que solo van a la capa de infra.

type UserRepository interface {
	GetUserByID(id int) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}