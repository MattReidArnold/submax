package service

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	submaxhttp "github.com/mattreidarnold/submax/internal/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

type service struct {
	echoRouter *echo.Echo
	port       int
}

func NewService(
	logger *slog.Logger,
	port int,
	contextPath string,
) *service {
	e := submaxhttp.NewRouter(
		contextPath,
		logger,
	)
	return &service{
		port:       port,
		echoRouter: e,
	}
}

func (s *service) Run() error {
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		err := s.echoRouter.Start(":" + strconv.Itoa(s.port))
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	return g.Wait()
}
