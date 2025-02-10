package mongo

import (
	"context"

	"github.com/reearth/server-scaffold/internal/infra/mongo/mongodoc"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	m *mongo.Database
}

var _ user.Repo = (*User)(nil)

func NewUser(db *mongo.Database) *User {
	return &User{m: db}
}

func (a *User) FindBySub(ctx context.Context, sub string) (*user.User, error) {
	res := a.col().FindOne(ctx, bson.M{"sub": sub})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var doc mongodoc.User
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.ToUser()
}

func (a *User) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	res := a.col().FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var doc mongodoc.User
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.ToUser()
}

func (a *User) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	res := a.col().FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var doc mongodoc.User
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.ToUser()
}

func (a *User) FindByIDs(ctx context.Context, ids user.IDList) (user.List, error) {
	res, err := a.col().Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.UserList
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToUserList()
}

func (a *User) FindByProject(ctx context.Context, pid project.ID) (user.List, error) {
	res, err := a.col().Find(ctx, bson.M{"project": pid})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.UserList
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToUserList()
}

func (a *User) Save(ctx context.Context, user *user.User) error {
	doc, err := mongodoc.NewUser(user)
	if err != nil {
		return err
	}

	err = a.col().FindOneAndUpdate(
		ctx,
		bson.M{"id": user.ID()},
		bson.M{"$set": doc},
		options.FindOneAndUpdate().SetUpsert(true),
	).Err()
	return err
}

func (a *User) Delete(ctx context.Context, id user.ID) error {
	_, err := a.col().DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (a *User) col() *mongo.Collection {
	return a.m.Collection("user")
}
