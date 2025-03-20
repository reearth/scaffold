package mongo

import (
	"context"

	"github.com/reearth/scaffold/server/internal/infra/mongo/mongodoc"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
	m *mongo.Database
}

var _ todo.Repo = (*Todo)(nil)

func NewTodo(db *mongo.Database) *Todo {
	return &Todo{m: db}
}

func (a *Todo) FindByID(ctx context.Context, id todo.ID) (*todo.Todo, error) {
	res := a.col().FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var doc mongodoc.Todo
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.Into()
}

func (a *Todo) FindByIDs(ctx context.Context, ids todo.IDList) (todo.List, error) {
	res, err := a.col().Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.List
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.Into()
}

func (a *Todo) FindByProject(ctx context.Context, pid project.ID) (todo.List, error) {
	res, err := a.col().Find(ctx, bson.M{"project": pid})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.List
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.Into()
}

func (a *Todo) Save(ctx context.Context, asset *todo.Todo) error {
	doc, err := mongodoc.New(asset)
	if err != nil {
		return err
	}

	err = a.col().FindOneAndUpdate(
		ctx,
		bson.M{"id": asset.ID()},
		bson.M{"$set": doc},
		options.FindOneAndUpdate().SetUpsert(true),
	).Err()
	return err
}

func (a *Todo) Delete(ctx context.Context, id todo.ID) error {
	_, err := a.col().DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (a *Todo) col() *mongo.Collection {
	return a.m.Collection("todo")
}
