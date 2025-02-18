package project

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/derror"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

type Policy interface {
	CanRead(ctx context.Context, workspace *workspace.Workspace, user *user.User, project *Project) error
	CanUpdate(ctx context.Context, workspace *workspace.Workspace, user *user.User, project *Project) error
	CanDelete(ctx context.Context, workspace *workspace.Workspace, user *user.User, project *Project) error
	Filter(ctx context.Context, workspaces workspace.List, user *user.User, project *Project) workspace.List
}

type DefaultPolicy struct{}

func NewPolicy() Policy {
	return DefaultPolicy{}
}

func (DefaultPolicy) CanRead(ctx context.Context, ws *workspace.Workspace, user *user.User, project *Project) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleMember) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (DefaultPolicy) CanUpdate(ctx context.Context, ws *workspace.Workspace, user *user.User, project *Project) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleAdmin) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (DefaultPolicy) CanDelete(ctx context.Context, ws *workspace.Workspace, user *user.User, project *Project) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleAdmin) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (DefaultPolicy) Filter(ctx context.Context, workspaces workspace.List, user *user.User, project *Project) workspace.List {
	res := make(workspace.List, 0, len(workspaces))
	for _, ws := range workspaces {
		if ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleMember) {
			res = append(res, ws)
		}
	}
	return res
}
