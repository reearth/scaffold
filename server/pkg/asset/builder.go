package asset

import (
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/samber/lo"
)

type Builder struct {
	a *Asset
}

func New() *Builder {
	return &Builder{a: &Asset{}}
}

func (b *Builder) Build() (*Asset, error) {
	if err := b.a.Validate(); err != nil {
		return nil, err
	}
	return b.a, nil
}

func (b *Builder) MustBuild() *Asset {
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
