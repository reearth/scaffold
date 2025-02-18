package user

import (
	"github.com/samber/lo"
)

type Builder struct {
	u *User
}

func New() *Builder {
	return &Builder{u: &User{}}
}

func (b *Builder) Build() (*User, error) {
	if err := b.u.Validate(); err != nil {
		return nil, err
	}
	return b.u, nil
}

func (b *Builder) MustBuild() *User {
	return lo.Must(b.Build())
}

func (b *Builder) ID(id ID) *Builder {
	b.u.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.u.id = NewID()
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.u.name = name
	return b
}

func (b *Builder) Email(email string) *Builder {
	b.u.email = email
	return b
}
