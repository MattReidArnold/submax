package http

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	contextPath string,
	logger *slog.Logger,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(LoggerMiddleware(logger))
	e.Use(middleware.Recover())

	g := e.Group(contextPath)

	g.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	g.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "world")
	})
	return e
}
