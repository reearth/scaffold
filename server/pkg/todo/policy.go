package todo

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/derror"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

type Policy interface {
	CanRead(ctx context.Context, user *user.User, workspace *workspace.Workspace, project *project.Project, asset *Todo) error
	CanList(ctx context.Context, user *user.User, workspace *workspace.Workspace, project *project.Project) error
	CanCreate(ctx context.Context, user *user.User, workspace *workspace.Workspace, project *project.Project) error
	CanUpdate(ctx context.Context, user *user.User, workspace *workspace.Workspace, project *project.Project, asset *Todo) error
	CanDelete(ctx context.Context, user *user.User, workspace *workspace.Workspace, project *project.Project, asset *Todo) error
	Filter(ctx context.Context, user *user.User, workspaces workspace.List, projects project.List, assets List) (List, error)
}

type DefaultPolicy struct{}

func NewPolicy() Policy {
	return DefaultPolicy{}
}

func (DefaultPolicy) CanRead(ctx context.Context, user *user.User, ws *workspace.Workspace, project *project.Project, asset *Todo) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleMember) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (DefaultPolicy) CanList(ctx context.Context, user *user.User, ws *workspace.Workspace, project *project.Project) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleMember) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (DefaultPolicy) CanCreate(ctx context.Context, user *user.User, ws *workspace.Workspace, project *project.Project) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleMember) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (DefaultPolicy) CanUpdate(ctx context.Context, user *user.User, ws *workspace.Workspace, project *project.Project, asset *Todo) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleMember) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (DefaultPolicy) CanDelete(ctx context.Context, user *user.User, ws *workspace.Workspace, project *project.Project, asset *Todo) error {
	if !ws.Members().HasRoleOrHigher(user.ID(), workspace.RoleMember) {
		return derror.ErrPermissionDenied
	}
	return nil
}

func (p DefaultPolicy) Filter(ctx context.Context, user *user.User, workspaces workspace.List, projects project.List, assets List) (List, error) {
	res := make(List, 0, len(assets))
	for _, a := range assets {
		prj := projects.Get(a.Project())
		if prj == nil {
			continue
		}

		ws := workspaces.Get(prj.Workspace())
		if ws == nil {
			continue
		}

		if err := p.CanRead(ctx, user, ws, prj, a); err != nil {
			continue
		}

		res = append(res, a)
	}
	return res, nil
}
