package workspace

import "context"

type Repo interface {
	FindByID(ctx context.Context, id ID) (*Workspace, error)
	FindByIDs(ctx context.Context, ids IDList) (List, error)
	Save(ctx context.Context, workspace *Workspace) error
	Delete(ctx context.Context, id ID) error
}
