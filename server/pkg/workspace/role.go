package workspace

type Role string

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
)

func (r Role) Compare(other Role) int {
	if r == other {
		return 0
	}
	if r == RoleOwner {
		return 1
	}
	if other == RoleOwner {
		return -1
	}
	if r == RoleAdmin {
		return 1
	}
	if other == RoleAdmin {
		return -1
	}
	return 0
}
