package server

import (
	"github.com/maiaaraujo5/go-project-template/internal/app/project/fx/service"
	"github.com/maiaaraujo5/go-project-template/internal/app/project/server/handler"
	"github.com/maiaaraujo5/go-project-template/internal/app/project/server/routes"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		GetterModule(),
		fx.Invoke(routes.Add),
	)
}

func GetterModule() fx.Option {
	return fx.Options(
		service.Module(),
		fx.Provide(
			fx.Annotate(
				handler.NewGetter,
				fx.ResultTags(`group:"handlers"`),
				fx.As(new(handler.Handle))),
		),
	)
}
