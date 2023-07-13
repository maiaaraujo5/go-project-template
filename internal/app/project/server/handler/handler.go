package handler

import "github.com/labstack/echo/v4"

type Handle interface {
	Handler(c echo.Context) error
	AddRoute(c *echo.Echo)
}
