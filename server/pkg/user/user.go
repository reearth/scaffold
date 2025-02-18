package user

import "errors"

type User struct {
	id    ID
	name  string
	email string
}

type List []*User

// getters

func (u *User) ID() ID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

// setters

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) error {
	// TODO: validate email
	u.email = email
	return nil
}

func (u *User) Validate() error {
	if u.id == (ID{}) {
		return errors.New("ID is required")
	}
	if u.name == "" {
		return errors.New("Name is required")
	}
	if u.email == "" {
		return errors.New("Email is required")
	}
	return nil
}

func (u *User) Clone() *User {
	if u == nil {
		return nil
	}
	return &User{
		id:    u.id,
		name:  u.name,
		email: u.email,
	}
}
