package domain

import (
	"errors"
	"regexp"
)

const (
	EMAIL_REGEX = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	NAME_REGEX  = `^[a-zA-Z]+$`
)

type User struct {
	ID       int
	Name     string
	LastName string
	Email    string
}

func (u User) ValitateNewUser() error {
	emailPattern := regexp.MustCompile(EMAIL_REGEX)
	namePattern := regexp.MustCompile(NAME_REGEX)

	if !emailPattern.MatchString(u.Email) {
		return errors.New("Invalid email")
	}

	if !namePattern.MatchString(u.Name) {
		return errors.New("Invalid name")
	}

	if !namePattern.MatchString(u.LastName) {
		return errors.New("Invalid last name")
	}

	return nil
}