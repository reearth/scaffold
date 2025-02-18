package asset

import (
	"github.com/reearth/scaffold/server/pkg/project"
)

type List []*Asset

func (l List) ProjectIDs() project.IDList {
	ids := make(project.IDList, 0, len(l))
	for _, a := range l {
		ids = append(ids, a.Project())
	}
	return ids
}
