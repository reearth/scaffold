package project

import (
	"github.com/samber/lo"
)

type ProjectBuilder struct {
	p *Project
}

func NewProjectBuilder() *ProjectBuilder {
	return &ProjectBuilder{&Project{}}
}

func (pb *ProjectBuilder) Build() (*Project, error) {
	if err := pb.p.Validate(); err != nil {
		return nil, err
	}
	return pb.p, nil
}

func (pb *ProjectBuilder) MustBuild() *Project {
	return lo.Must(pb.Build())
}

func (pb *ProjectBuilder) ID(id ID) *ProjectBuilder {
	pb.p.id = id
	return pb
}

func (pb *ProjectBuilder) Workspace(workspace ID) *ProjectBuilder {
	pb.p.workspace = workspace
	return pb
}

func (pb *ProjectBuilder) Name(name string) *ProjectBuilder {
	pb.p.name = name
	return pb
}
