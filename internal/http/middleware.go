package http

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			if err := next(c); err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			latency := time.Since(start)
			msg := fmt.Sprintf(
				"Incoming request: %s %s status=%d latency=%s",
				req.Method,
				req.RequestURI,
				res.Status,
				latency,
			)

			logger.InfoContext(
				c.Request().Context(),
				msg,
				slog.String("method", req.Method),
				slog.String("uri", req.RequestURI),
				slog.String("ip", c.RealIP()),
				slog.String("user-agent", req.UserAgent()),
				slog.Int("status", res.Status),
				slog.Duration("latency", latency),
			)

			return nil
		}
	}
}
