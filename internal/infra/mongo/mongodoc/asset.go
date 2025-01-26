package mongodoc

import (
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
)

type Asset struct {
	ID      asset.ID   `bson:"id"`
	Project project.ID `bson:"project"`
	Name    string     `bson:"name"`
}

func (a *Asset) ToAsset() (*asset.Asset, error) {
	return asset.New().
		ID(a.ID).
		Project(a.Project).
		Name(a.Name).
		Build()
}

func New(a *asset.Asset) (*Asset, error) {
	return &Asset{
		ID:      a.ID(),
		Project: a.Project(),
		Name:    a.Name(),
	}, nil
}

type List []Asset

func (l List) ToAssetList() (asset.List, error) {
	assets := make(asset.List, 0, len(l))
	for _, a := range l {
		asset, err := a.ToAsset()
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func NewList(assets asset.List) (List, error) {
	list := make(List, 0, len(assets))
	for _, a := range assets {
		doc, err := New(a)
		if err != nil {
			return nil, err
		}
		list = append(list, *doc)
	}
	return list, nil
}
