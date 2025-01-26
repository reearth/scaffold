package project

import (
	"errors"

	"github.com/reearth/server-scaffold/pkg/workspace"
	"github.com/samber/lo"
)

type Project struct {
	id        ID
	workspace workspace.ID
	name      string
}

// getters

func (p *Project) ID() ID {
	return p.id
}

func (p *Project) Workspace() workspace.ID {
	return p.workspace
}

func (p *Project) Name() string {
	return p.name
}

// setters

func (p *Project) SetName(name string) {
	p.name = name
}

func (p *Project) Validate() error {
	if lo.IsEmpty(p.id) {
		return errors.New("id is required")
	}
	if lo.IsEmpty(p.workspace) {
		return errors.New("workspace is required")
	}
	if p.name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (p *Project) Clone() *Project {
	if p == nil {
		return nil
	}
	return &Project{
		id:        p.id,
		workspace: p.workspace,
		name:      p.name,
	}
}
