package handler

import "github.com/labstack/echo/v4"

type healthHandler struct{}

func NewHealthHandler() Handler {
	return healthHandler{}
}

func (h healthHandler) Register(e *echo.Echo) {
	e.GET("/health", h.Health)
}

func (h healthHandler) Health(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "ok",
	})
}
