package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/go-project-template/internal/app/project/server/handler"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handlers []handler.Handle `group:"handlers"`
	Echo     *echo.Echo
}
