package gateway

import (
	"context"
)

type Transaction interface {
	Begin() Tx
	Do(context.Context, func() error) error
}

type Tx interface {
	Commit(context.Context) error
	Rollback(context.Context) error
	CommitOrRollback(context.Context, *error)
}

type NoopTransaction struct{}

func (t *NoopTransaction) Begin() Tx {
	return &NoopTx{}
}

type NoopTx struct {
}

func (t *NoopTx) Commit(context.Context) error {
	return nil
}

func (t *NoopTx) Rollback(context.Context) error {
	return nil
}

func (t *NoopTx) CommitOrRollback(context.Context, *error) {
}
