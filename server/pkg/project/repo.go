package project

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/workspace"
)

type Repo interface {
	FindByID(ctx context.Context, id ID) (*Project, error)
	FindByIDs(ctx context.Context, id IDList) (List, error)
	FindByWorkspaceID(ctx context.Context, workspaceID workspace.ID) (List, error)
	Save(ctx context.Context, project *Project) error
	Delete(ctx context.Context, id ID) error
}
