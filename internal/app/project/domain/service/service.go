package service

import (
	"context"

	"github.com/maiaaraujo5/go-project-template/internal/app/project/domain/repository"
)

type Finder interface {
	Execute(ctx context.Context) error
}

type Find struct {
	repository repository.Finder
}

func NewService(repository repository.Finder) *Find {
	return &Find{
		repository: repository,
	}
}

func (s *Find) Execute(ctx context.Context) error {
	err := s.repository.Find(ctx)
	if err != nil {
		return err
	}

	return nil
}
