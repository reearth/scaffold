package user

import "context"

type Repo interface {
	FindByID(ctx context.Context, id ID) (*User, error)
	FindByIDs(ctx context.Context, ids IDList) (List, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindBySub(ctx context.Context, sub string) (*User, error)
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id ID) error
}
