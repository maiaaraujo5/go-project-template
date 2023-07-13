package service

import (
	"context"
	"github.com/maiaaraujo5/go-project-template/internal/app/project/domain/repository"
)

type service interface {
	Execute(ctx context.Context) error
}

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Execute(ctx context.Context) error {
	err := s.repository.Find(ctx)
	if err != nil {
		return err
	}

	return nil
}
