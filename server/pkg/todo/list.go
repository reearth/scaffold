package todo

import (
	"github.com/reearth/scaffold/server/pkg/project"
)

type List []*Todo

func (l List) ProjectIDs() project.IDList {
	ids := make(project.IDList, 0, len(l))
	for _, a := range l {
		ids = append(ids, a.Project())
	}
	return ids
}
