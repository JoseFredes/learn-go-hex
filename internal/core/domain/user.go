package domain

import "regexp"

const (
	EMAIL_REGEX = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type User struct {
	ID int
	Name string
	LastName string
	Email string
}


func (u *User) UpdateName(name string) string {
	u.Name = name
	return u.Name
}

func (u *User) UpdateLastName(lastName string) string {
	u.LastName = lastName
	return u.LastName
}

func (u *User) UpdateEmail(email string) string {
	u.Email = email
	return u.Email
}

func (u *User) ValidateEmail() bool {
	pattern := regexp.MustCompile(EMAIL_REGEX)
	return pattern.MatchString(u.Email)
}

func NewUser(id int, name string, lastName string, email string) User {
	return User{ID: id, Name: name, LastName: lastName, Email: email}
}