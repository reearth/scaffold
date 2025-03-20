package todo

import (
	"time"

	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/samber/lo"
)

type Builder struct {
	a *Todo
}

func New() *Builder {
	return &Builder{a: &Todo{}}
}

func (b *Builder) Build() (*Todo, error) {
	if err := b.a.Validate(); err != nil {
		return nil, err
	}
	return b.a, nil
}

func (b *Builder) MustBuild() *Todo {
	return lo.Must(b.Build())
}

func (b *Builder) ID(id ID) *Builder {
	b.a.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.a.id = NewID()
	return b
}

func (b *Builder) Project(project project.ID) *Builder {
	b.a.project = project
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.a.name = name
	return b
}

func (b *Builder) Done(done bool) *Builder {
	b.a.done = done
	return b
}

func (b *Builder) CreatedAt(createdAt time.Time) *Builder {
	b.a.createdAt = createdAt
	return b
}

func (b *Builder) UpdatedAt(updatedAt time.Time) *Builder {
	b.a.updatedAt = updatedAt
	return b
}
