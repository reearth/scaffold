package todo

import (
	"errors"
	"time"

	"github.com/reearth/scaffold/server/pkg/project"
)

// types

type Todo struct {
	id        ID
	project   project.ID
	name      string
	done      bool
	createdAt time.Time
	updatedAt time.Time
}

// getters

func (a *Todo) ID() ID {
	return a.id
}

func (a *Todo) Project() project.ID {
	return a.project
}

func (a *Todo) Name() string {
	return a.name
}

func (a *Todo) Done() bool {
	return a.done
}

// setters

func (a *Todo) SetName(name string) {
	a.name = name
}

func (a *Todo) Validate() error {
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

func (a *Todo) Clone() *Todo {
	if a == nil {
		return nil
	}
	return &Todo{
		id:        a.id,
		name:      a.name,
		done:      a.done,
		createdAt: a.createdAt,
		updatedAt: a.updatedAt,
	}
}
