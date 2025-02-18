package workspace

import (
	"fmt"

	"github.com/reearth/scaffold/server/pkg/user"
)

type Members struct {
	m map[user.ID]Member
}

func NewMembers(m map[user.ID]Member) (Members, error) {
	members := Members{m: copyMembersMap(m)}
	if err := members.Validate(); err != nil {
		return Members{}, err
	}
	return members, nil
}

func (m Members) Add(member Member) {
	if m.m == nil {
		m.m = make(map[user.ID]Member)
	}
	m.m[member.User()] = member
}

func (m Members) Remove(id user.ID) {
	delete(m.m, id)
}

func (m Members) Get(id user.ID) *Member {
	if m.m == nil {
		return nil
	}
	if v, ok := m.m[id]; ok {
		return &v
	}
	return nil
}

func (m Members) All() map[user.ID]Member {
	if m.m == nil {
		return nil
	}
	return copyMembersMap(m.m)
}

func (m Members) HasRole(id user.ID, role Role) bool {
	if m.m == nil {
		return false
	}
	if v, ok := m.m[id]; ok {
		if v.Role().Compare(role) == 0 {
			return true
		}
	}
	return false
}

func (m Members) HasRoleOrHigher(id user.ID, role Role) bool {
	if m.m == nil {
		return false
	}
	if v, ok := m.m[id]; ok {
		if v.Role().Compare(role) >= 0 {
			return true
		}
	}
	return false
}

func (m *Members) Validate() error {
	for k, v := range m.m {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("member %s: %w", k, err)
		}
		if k != v.User() {
			return fmt.Errorf("member %s: key and value user id mismatch", k)
		}
	}
	return nil
}

func (m *Members) Clone() Members {
	if m == nil {
		return Members{}
	}
	n := make(map[user.ID]Member, len(m.m))
	for k, v := range m.m {
		n[k] = *v.Clone()
	}
	return Members{n}
}

func copyMembersMap(m map[user.ID]Member) map[user.ID]Member {
	if m == nil {
		return nil
	}
	n := make(map[user.ID]Member, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}
