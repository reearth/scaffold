package workspace

import (
	"fmt"

	"github.com/samber/lo"
)

type Workspace struct {
	id      ID
	members Members
}

func (w *Workspace) ID() ID {
	return w.id
}

func (w *Workspace) Members() Members {
	return w.members
}

func (v *Workspace) Validate() error {
	if lo.IsEmpty(v.id) {
		return fmt.Errorf("workspace id is required")
	}
	if err := v.members.Validate(); err != nil {
		return err
	}
	return nil
}
