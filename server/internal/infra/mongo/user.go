package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/reearth/scaffold/server/internal/infra/mongo/mongodoc"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	m *mongo.Database
}

var _ user.Repo = (*User)(nil)

var ErrUserNotFound = errors.New("user not found")

func NewUser(db *mongo.Database) *User {
	return &User{m: db}
}

func (a *User) FindBySub(ctx context.Context, sub string) (*user.User, error) {
	return a.findOne(ctx, bson.M{"sub": sub})
}

func (a *User) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	return a.findOne(ctx, bson.M{"email": email})
}

func (a *User) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	return a.findOne(ctx, bson.M{"id": id})
}

func (a *User) FindByIDs(ctx context.Context, ids user.IDList) (user.List, error) {
	return a.findMany(ctx, bson.M{"id": bson.M{"$in": ids}})
}

func (a *User) FindByProject(ctx context.Context, pid project.ID) (user.List, error) {
	return a.findMany(ctx, bson.M{"project": pid})
}

func (a *User) Save(ctx context.Context, user *user.User) error {
	if user == nil {
		return fmt.Errorf("user is required")
	}

	doc, err := mongodoc.NewUser(user)
	if err != nil {
		return err
	}

	filter := bson.M{"id": user.ID()}
	update := bson.M{"$set": doc}

	err = a.col().FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetUpsert(true),
	).Err()

	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("concurrent modification detected")
	}

	return err
}

func (a *User) Delete(ctx context.Context, id user.ID) error {
	_, err := a.col().DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (a *User) col() *mongo.Collection {
	return a.m.Collection("user")
}

func (a *User) findOne(ctx context.Context, filter bson.M) (*user.User, error) {
	var doc mongodoc.User
	if err := a.col().FindOne(ctx, filter).Decode(&doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return doc.ToUser()
}

func (a *User) findMany(ctx context.Context, filter bson.M, opts ...*options.FindOptions) (user.List, error) {
	var docs mongodoc.UserList
	cursor, err := a.col().Find(ctx, filter, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("failed to decode users: %w", err)
	}
	return docs.ToUserList()
}
