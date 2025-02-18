package workspace

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/derror"
	"github.com/reearth/server-scaffold/pkg/user"
)

type Policy interface {
	CanRead(ctx context.Context, workspace *Workspace, user *user.User) error
	CanCreate(ctx context.Context, user *user.User) error
	CanUpdate(ctx context.Context, workspace *Workspace, user *user.User) error
	CanDelete(ctx context.Context, workspace *Workspace, user *user.User) error
	Filter(ctx context.Context, workspaces List, user *user.User) List
}

type DefaultPolicy struct{}

func NewPolicy() Policy {
	return DefaultPolicy{}
}

func (p DefaultPolicy) CanRead(ctx context.Context, ws *Workspace, user *user.User) error {
	if !ws.members.HasRoleOrHigher(user.ID(), RoleMember) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (p DefaultPolicy) CanCreate(ctx context.Context, user *user.User) error {
	return nil // all users can create workspace
}

func (p DefaultPolicy) CanUpdate(ctx context.Context, ws *Workspace, user *user.User) error {
	if !ws.members.HasRoleOrHigher(user.ID(), RoleAdmin) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (p DefaultPolicy) CanDelete(ctx context.Context, ws *Workspace, user *user.User) error {
	if !ws.members.HasRoleOrHigher(user.ID(), RoleOwner) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (p DefaultPolicy) Filter(ctx context.Context, workspaces List, user *user.User) List {
	res := make(List, 0, len(workspaces))
	for _, ws := range workspaces {
		if ws.members.HasRoleOrHigher(user.ID(), RoleMember) {
			res = append(res, ws)
		}
	}
	return res
}
