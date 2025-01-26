package assetuc

import (
	"context"
	"errors"

	"github.com/reearth/server-scaffold/internal/usecase"
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

type Usecase struct {
	usecase.Deps
}

func New(uc usecase.Deps) *Usecase {
	return &Usecase{
		Deps: uc,
	}
}

type Builder struct {
	uc        *Usecase
	ctx       context.Context
	err       error
	user      *user.User
	asset     *asset.Asset
	project   *project.Project
	workspace *workspace.Workspace
}

func (uc *Usecase) Builder(ctx context.Context, user *user.User) *Builder {
	return &Builder{uc: uc, ctx: ctx, user: user}
}

func (b *Builder) Result() (*asset.Asset, *project.Project, *workspace.Workspace, error) {
	if b.err != nil {
		return nil, nil, nil, b.err
	}
	return b.asset, b.project, b.workspace, b.err
}

func (b *Builder) FindAssetByID(id asset.ID) *Builder {
	if b.err != nil {
		return b
	}
	b.asset, b.err = b.uc.Repos.Asset.FindByID(b.ctx, id)
	return b
}

func (b *Builder) FindProjectByID(id project.ID) *Builder {
	if b.err != nil {
		return b
	}
	b.project, b.err = b.uc.Repos.Project.FindByID(b.ctx, id)
	if b.err == nil {
		b.workspace, b.err = b.uc.Repos.Workspace.FindByID(b.ctx, b.project.Workspace())
	}
	return b
}

func (b *Builder) FindProjectByAsset() *Builder {
	if b.err != nil {
		return b
	}
	if b.asset == nil {
		b.err = errors.New("asset not found")
		return b
	}
	b.project, b.err = b.uc.Repos.Project.FindByID(b.ctx, b.asset.Project())
	if b.err == nil {
		b.workspace, b.err = b.uc.Repos.Workspace.FindByID(b.ctx, b.project.Workspace())
	}
	return b
}

func (b *Builder) CanReadAsset() *Builder {
	if b.err != nil {
		return b
	}
	b.err = b.uc.Policies.Asset.CanRead(b.ctx, b.user, b.workspace, b.project, b.asset)
	return b
}

func (b *Builder) CanListAssets() *Builder {
	if b.err != nil {
		return b
	}
	b.err = b.uc.Policies.Asset.CanList(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *Builder) CanCreateAsset() *Builder {
	if b.err != nil {
		return b
	}
	b.err = b.uc.Policies.Asset.CanCreate(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *Builder) CanUpdateAsset() *Builder {
	if b.err != nil {
		return b
	}
	b.err = b.uc.Policies.Asset.CanUpdate(b.ctx, b.user, b.workspace, b.project, b.asset)
	return b
}

func (b *Builder) CanDeleteAsset() *Builder {
	if b.err != nil {
		return b
	}
	b.err = b.uc.Policies.Asset.CanDelete(b.ctx, b.user, b.workspace, b.project, b.asset)
	return b
}
