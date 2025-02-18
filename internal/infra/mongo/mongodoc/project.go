package mongodoc

import (
	"github.com/reearth/server-scaffold/pkg/project"
)

type Project struct {
	ID      project.ID `bson:"id"`
	Project project.ID `bson:"project"`
	Name    string     `bson:"name"`
}

func (a *Project) ToProject() (*project.Project, error) {
	return &project.Project{}, nil
}

func NewProject(a *project.Project) (*Project, error) {
	return &Project{
		ID: a.ID(),
	}, nil
}

type ProjectList []Project

func (l ProjectList) ToProjectList() (project.List, error) {
	projects := make(project.List, 0, len(l))
	for _, a := range l {
		project, err := a.ToProject()
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func NewProjectList(projects project.List) (ProjectList, error) {
	list := make(ProjectList, 0, len(projects))
	for _, a := range projects {
		doc, err := NewProject(a)
		if err != nil {
			return nil, err
		}
		list = append(list, *doc)
	}
	return list, nil
}
