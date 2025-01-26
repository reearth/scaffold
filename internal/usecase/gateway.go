package usecase

import (
	"context"
	"io"
)

type Gateways struct {
	Storage Storage
}

type Storage interface {
	Open(ctx context.Context, name string) (io.ReadCloser, error)
	Save(ctx context.Context, name string, data io.Reader) error
	Delete(ctx context.Context, name string) error
}
