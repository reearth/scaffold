package workspace

import (
	"fmt"

	"github.com/samber/lo"
)

type Builder struct {
	w *Workspace
}

func New() *Builder {
	return &Builder{w: &Workspace{}}
}

func (b *Builder) Build() (*Workspace, error) {
	if lo.IsEmpty(b.w.id) {
		return nil, fmt.Errorf("workspace id is required")
	}
	return b.w, nil
}

func (b *Builder) MustBuild() *Workspace {
	return lo.Must(b.Build())
}

func (b *Builder) ID(id ID) *Builder {
	b.w.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.w.id = NewID()
	return b
}

func (b *Builder) Members(m Members) *Builder {
	b.w.members = m
	return b
}
