package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/go-project-template/internal/app/project/domain/service"
)

type Getter struct {
	service service.Finder
}

func NewGetter(service service.Finder) *Getter {
	return &Getter{
		service: service,
	}
}

func (h *Getter) AddRoute(c *echo.Echo) {
	c.GET("/", h.Handler)
}

// Handler godoc
// @Summary Find
// @Description Find
// @Tags user
// @Accept json
// @Produce json
// @Router /v1/ [get]
func (h *Getter) Handler(c echo.Context) error {
	ctx := c.Request().Context()

	err := h.service.Execute(ctx)
	if err != nil {
		return err
	}

	return nil
}
