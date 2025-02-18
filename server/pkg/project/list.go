package project

import "github.com/reearth/scaffold/server/pkg/workspace"

type List []*Project

func (l List) WorkspaceIDs() workspace.IDList {
	ids := make(workspace.IDList, 0, len(l))
	for _, p := range l {
		ids = append(ids, p.Workspace())
	}
	return ids
}

func (l List) Get(id ID) *Project {
	for _, p := range l {
		if p.ID() == id {
			return p
		}
	}
	return nil
}
