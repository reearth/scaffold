package mongodoc

import (
	"github.com/reearth/scaffold/server/pkg/user"
)

type User struct {
	ID   user.ID `bson:"id"`
	Name string  `bson:"name"`
}

func (a *User) ToUser() (*user.User, error) {
	return user.New().
		ID(a.ID).
		Build()
}

func NewUser(a *user.User) (*User, error) {
	return &User{
		ID:   a.ID(),
		Name: a.Name(),
	}, nil
}

type UserList []User

func (l UserList) ToUserList() (user.List, error) {
	users := make(user.List, 0, len(l))
	for _, a := range l {
		user, err := a.ToUser()
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func NewUserList(users user.List) (UserList, error) {
	list := make(UserList, 0, len(users))
	for _, a := range users {
		doc, err := NewUser(a)
		if err != nil {
			return nil, err
		}
		list = append(list, *doc)
	}
	return list, nil
}
