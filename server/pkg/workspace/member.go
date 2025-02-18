package workspace

import (
	"errors"

	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/samber/lo"
)

type Member struct {
	user user.ID
	role Role
}

func (m *Member) User() user.ID {
	return m.user
}

func (m *Member) Role() Role {
	return m.role
}

func (m *Member) SetRole(role Role) {
	m.role = role
}

func (m *Member) Validate() error {
	if lo.IsEmpty(m.user) {
		return errors.New("user is required")
	}
	if m.role == "" {
		return errors.New("role is required")
	}
	return nil
}

func (m *Member) Clone() *Member {
	if m == nil {
		return nil
	}
	return &Member{
		user: m.user,
		role: m.role,
	}
}
