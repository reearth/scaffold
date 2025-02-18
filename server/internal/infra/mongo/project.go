package mongo

import (
	"context"

	"github.com/reearth/scaffold/server/internal/infra/mongo/mongodoc"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/workspace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Project struct {
	m *mongo.Database
}

var _ project.Repo = (*Project)(nil)

func NewProject(db *mongo.Database) *Project {
	return &Project{m: db}
}

func (a *Project) FindByID(ctx context.Context, id project.ID) (*project.Project, error) {
	res := a.col().FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var doc mongodoc.Project
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.ToProject()
}

func (a *Project) FindByWorkspaceID(ctx context.Context, id workspace.ID) (project.List, error) {
	res, err := a.col().Find(ctx, bson.M{"workspace": id})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.ProjectList
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToProjectList()
}

func (a *Project) FindByIDs(ctx context.Context, ids project.IDList) (project.List, error) {
	res, err := a.col().Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.ProjectList
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToProjectList()
}

func (a *Project) FindByProject(ctx context.Context, pid project.ID) (project.List, error) {
	res, err := a.col().Find(ctx, bson.M{"project": pid})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.ProjectList
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToProjectList()
}

func (a *Project) Save(ctx context.Context, project *project.Project) error {
	doc, err := mongodoc.NewProject(project)
	if err != nil {
		return err
	}

	err = a.col().FindOneAndUpdate(
		ctx,
		bson.M{"id": project.ID()},
		bson.M{"$set": doc},
		options.FindOneAndUpdate().SetUpsert(true),
	).Err()
	return err
}

func (a *Project) Delete(ctx context.Context, id project.ID) error {
	_, err := a.col().DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (a *Project) col() *mongo.Collection {
	return a.m.Collection("project")
}
