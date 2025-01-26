package asset

import (
	"errors"

	"github.com/reearth/server-scaffold/pkg/project"
)

// types

type Asset struct {
	id      ID
	project project.ID
	name    string
}

// getters

func (a *Asset) ID() ID {
	return a.id
}

func (a *Asset) Project() project.ID {
	return a.project
}

func (a *Asset) Name() string {
	return a.name
}

// setters

func (a *Asset) SetName(name string) {
	a.name = name
}

func (a *Asset) Validate() error {
	if a.id == (ID{}) {
		return errors.New("ID is required")
	}
	if a.project == (project.ID{}) {
		return errors.New("Project is required")
	}
	if a.name == "" {
		return errors.New("Name is required")
	}
	return nil
}

func (a *Asset) Clone() *Asset {
	if a == nil {
		return nil
	}
	return &Asset{
		id:   a.id,
		name: a.name,
	}
}
