package gqlmodel

import (
	"github.com/reearth/scaffold/server/pkg/asset"
	"github.com/reearth/scaffold/server/pkg/project"
)

func NewAsset(a *asset.Asset) *Asset {
	if a == nil {
		return nil
	}
	return &Asset{
		ID:        ID(a.ID()),
		Name:      a.Name(),
		ProjectID: ID(a.Project()),
	}
}

func NewAssets(assets asset.List) []*Asset {
	if assets == nil {
		return nil
	}
	res := make([]*Asset, 0, len(assets))
	for _, a := range assets {
		res = append(res, NewAsset(a))
	}
	return res
}

func (a *Asset) Into() (*asset.Asset, error) {
	return asset.New().
		ID(asset.ID(a.ID)).
		Name(a.Name).
		Project(project.ID(a.ProjectID)).
		Build()
}
