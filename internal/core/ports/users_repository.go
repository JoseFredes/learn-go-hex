package ports

import "github.com/JoseFredes/learn-go-hex/internal/core/domain"

type UserServices interface {
	Get(id int) (domain.User, error)
}

type UserRepository interface {
	GetUserByID(id int) (domain.User, error)
}