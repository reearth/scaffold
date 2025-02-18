package workspace

import (
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/samber/lo"
)

type MemberBuilder struct {
	m *Member
}

func NewMember() *MemberBuilder {
	return &MemberBuilder{m: &Member{}}
}

func (b *MemberBuilder) Build() (*Member, error) {
	if err := b.m.Validate(); err != nil {
		return nil, err
	}
	return b.m, nil
}

func (b *MemberBuilder) MustBuild() *Member {
	return lo.Must(b.Build())
}

func (b *MemberBuilder) User(user user.ID) *MemberBuilder {
	b.m.user = user
	return b
}
