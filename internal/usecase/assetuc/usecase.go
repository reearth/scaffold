package assetuc

import (
	"context"
	"errors"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

type Usecase struct {
	FindByIDs     *FindByIDs
	FindByProject *FindByProject
	Create        *Create
	Update        *Update
}

func New(
	findByIDs *FindByIDs,
	findByProject *FindByProject,
	create *Create,
	update *Update,
) *Usecase {
	return &Usecase{
		FindByIDs:     findByIDs,
		FindByProject: findByProject,
		Create:        create,
		Update:        update,
	}
}

type Builder struct {
	ctx       context.Context
	err       error
	user      *user.User
	asset     *asset.Asset
	project   *project.Project
	workspace *workspace.Workspace
}

func UsecaseBuilder(ctx context.Context, user *user.User) *Builder {
	return &Builder{ctx: ctx, user: user}
}

func (b *Builder) Result() (*asset.Asset, *project.Project, *workspace.Workspace, error) {
	if b.err != nil {
		return nil, nil, nil, b.err
	}
	return b.asset, b.project, b.workspace, b.err
}

func (b *Builder) FindAssetByID(id asset.ID, assetRepo asset.Repo) *Builder {
	if b.err != nil {
		return b
	}
	b.asset, b.err = assetRepo.FindByID(b.ctx, id)
	return b
}

func (b *Builder) FindProjectByID(id project.ID, projectRepo project.Repo, workspaceRepo workspace.Repo) *Builder {
	if b.err != nil {
		return b
	}
	b.project, b.err = projectRepo.FindByID(b.ctx, id)
	if b.err == nil {
		b.workspace, b.err = workspaceRepo.FindByID(b.ctx, b.project.Workspace())
	}
	return b
}

func (b *Builder) FindProjectByAsset(projectRepo project.Repo, workspaceRepo workspace.Repo) *Builder {
	if b.err != nil {
		return b
	}
	if b.asset == nil {
		b.err = errors.New("asset not found")
		return b
	}
	b.project, b.err = projectRepo.FindByID(b.ctx, b.asset.Project())
	if b.err == nil {
		b.workspace, b.err = workspaceRepo.FindByID(b.ctx, b.project.Workspace())
	}
	return b
}

func (b *Builder) CanReadAsset(assetPolicy asset.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanRead(b.ctx, b.user, b.workspace, b.project, b.asset)
	return b
}

func (b *Builder) CanListAssets(assetPolicy asset.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanList(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *Builder) CanCreateAsset(assetPolicy asset.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanCreate(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *Builder) CanUpdateAsset(assetPolicy asset.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanUpdate(b.ctx, b.user, b.workspace, b.project, b.asset)
	return b
}

func (b *Builder) CanDeleteAsset(assetPolicy asset.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanDelete(b.ctx, b.user, b.workspace, b.project, b.asset)
	return b
}
