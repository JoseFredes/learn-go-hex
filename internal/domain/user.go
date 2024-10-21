package domain

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

