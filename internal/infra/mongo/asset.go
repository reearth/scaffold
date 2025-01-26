package mongo

import (
	"context"

	"github.com/reearth/server-scaffold/internal/infra/mongo/mongodoc"
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Asset struct {
	m *mongo.Database
}

var _ asset.Repo = (*Asset)(nil)

func NewAsset(db *mongo.Database) *Asset {
	return &Asset{m: db}
}

func (a *Asset) FindByID(ctx context.Context, id asset.ID) (*asset.Asset, error) {
	res := a.col().FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var doc mongodoc.Asset
	if err := res.Decode(&doc); err != nil {
		return nil, err
	}

	return doc.ToAsset()
}

func (a *Asset) FindByIDs(ctx context.Context, ids asset.IDList) (asset.List, error) {
	res, err := a.col().Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.List
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToAssetList()
}

func (a *Asset) FindByProject(ctx context.Context, pid project.ID) (asset.List, error) {
	res, err := a.col().Find(ctx, bson.M{"project": pid})
	if err != nil {
		return nil, err
	}

	var docs mongodoc.List
	if err := res.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToAssetList()
}

func (a *Asset) Save(ctx context.Context, asset *asset.Asset) error {
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

func (a *Asset) Delete(ctx context.Context, id asset.ID) error {
	_, err := a.col().DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (a *Asset) col() *mongo.Collection {
	return a.m.Collection("asset")
}
