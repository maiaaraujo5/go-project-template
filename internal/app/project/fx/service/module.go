package service

import (
	"github.com/maiaaraujo5/go-project-template/internal/app/project/domain/service"
	"github.com/maiaaraujo5/go-project-template/internal/app/project/fx/provider"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		provider.Module(),
		fx.Provide(
			fx.Annotate(service.NewService, fx.As(new(service.Service))),
		),
	)
}
