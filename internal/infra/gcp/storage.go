package gcp

import (
	"context"
	"io"

	"github.com/reearth/server-scaffold/internal/usecase"
)

type Storage struct{}

var _ usecase.Storage = (*Storage)(nil)

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Open(ctx context.Context, name string) (io.ReadCloser, error) {
	// TODO: implement
	return nil, nil
}

func (s *Storage) Save(ctx context.Context, name string, data io.Reader) error {
	// TODO: implement
	return nil
}

func (s *Storage) Delete(ctx context.Context, name string) error {
	// TODO: implement
	return nil
}
