package provider

import (
	"github.com/maiaaraujo5/go-project-template/internal/app/project/domain/repository"
	"github.com/maiaaraujo5/go-project-template/internal/app/project/provider"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				provider.NewProvider,
				fx.As(new(repository.Repository)),
			),
		),
	)
}
