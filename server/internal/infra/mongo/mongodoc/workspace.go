package mongodoc

import (
	"github.com/reearth/scaffold/server/pkg/workspace"
)

type Workspace struct {
	ID      workspace.ID      `bson:"id"`
	Members workspace.Members `bson:"members"`
}

func (a *Workspace) ToWorkspace() (*workspace.Workspace, error) {
	return workspace.New().
		ID(a.ID).
		Members(a.Members).
		Build()
}

func NewWorkspace(a *workspace.Workspace) (*Workspace, error) {
	return &Workspace{
		ID:      a.ID(),
		Members: a.Members(),
	}, nil
}

type WorkspaceList []Workspace

func (l WorkspaceList) ToWorkspaceList() (workspace.List, error) {
	workspaces := make(workspace.List, 0, len(l))
	for _, a := range l {
		workspace, err := a.ToWorkspace()
		if err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}
	return workspaces, nil
}

func NewWorkspaceList(workspaces workspace.List) (WorkspaceList, error) {
	list := make(WorkspaceList, 0, len(workspaces))
	for _, a := range workspaces {
		doc, err := NewWorkspace(a)
		if err != nil {
			return nil, err
		}
		list = append(list, *doc)
	}
	return list, nil
}
