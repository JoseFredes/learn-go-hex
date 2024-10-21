package ports

import "github.com/JoseFredes/learn-go-hex/internal/core/domain"

type UsersServices interface {
	Get(id int) (domain.User, error)
}