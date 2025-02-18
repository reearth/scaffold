package asset

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/project"
)

type Repo interface {
	FindByID(ctx context.Context, id ID) (*Asset, error)
	FindByIDs(ctx context.Context, ids IDList) (List, error)
	FindByProject(ctx context.Context, projectID project.ID) (List, error)
	Save(ctx context.Context, a *Asset) error
	Delete(ctx context.Context, id ID) error
}
