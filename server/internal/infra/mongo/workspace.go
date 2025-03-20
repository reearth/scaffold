package mongo

import (
	"context"

	"github.com/reearth/scaffold/server/internal/infra/mongo/mongodoc"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/workspace"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Workspace struct {
	m *mongo.Database
}

var _ workspace.Repo = (*Workspace)(nil)

func NewWorkspace(db *mongo.Database) *Workspace {
	return &Workspace{m: db}
}

func (a *Workspace) FindByID(ctx context.Context, id workspace.ID) (*workspace.Workspace, error) {
	res := a.col().FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var doc mongodoc.Workspace
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.ToWorkspace()
}

func (a *Workspace) FindByIDs(ctx context.Context, ids workspace.IDList) (workspace.List, error) {
	res, err := a.col().Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.WorkspaceList
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToWorkspaceList()
}

func (a *Workspace) FindByProject(ctx context.Context, pid project.ID) (workspace.List, error) {
	res, err := a.col().Find(ctx, bson.M{"project": pid})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.WorkspaceList
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToWorkspaceList()
}

func (a *Workspace) Save(ctx context.Context, workspace *workspace.Workspace) error {
	doc, err := mongodoc.NewWorkspace(workspace)
	if err != nil {
		return err
	}

	err = a.col().FindOneAndUpdate(
		ctx,
		bson.M{"id": workspace.ID()},
		bson.M{"$set": doc},
		options.FindOneAndUpdate().SetUpsert(true),
	).Err()
	return err
}

func (a *Workspace) Delete(ctx context.Context, id workspace.ID) error {
	_, err := a.col().DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (a *Workspace) col() *mongo.Collection {
	return a.m.Collection("workspace")
}
