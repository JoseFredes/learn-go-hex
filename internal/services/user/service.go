package user

import "github.com/JoseFredes/learn-go-hex/internal/core/domain"

type userServices struct {}

func NewUserServices() *userServices {
	return &userServices{}
}


func (s *userServices) Get(id string, a int) (domain.User , error){
	return domain.User{} , nil
}