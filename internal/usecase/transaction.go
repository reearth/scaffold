package usecase

import (
	"context"
)

type Transaction interface {
	Begin() Tx
	Do(context.Context, func(Repos) error) error
}

type Tx interface {
	Get() Repos
	Commit(context.Context) error
	Rollback(context.Context) error
	CommitOrRollback(context.Context, *error)
}

type NoopTransaction struct {
	Repos Repos
}

func (t *NoopTransaction) Begin() Tx {
	return &NoopTx{Repos: t.Repos}
}

type NoopTx struct {
	Repos Repos
}

func (t *NoopTx) Get() Repos {
	return t.Repos
}

func (t *NoopTx) Commit(context.Context) error {
	return nil
}

func (t *NoopTx) Rollback(context.Context) error {
	return nil
}

func (t *NoopTx) CommitOrRollback(context.Context, *error) {
}
