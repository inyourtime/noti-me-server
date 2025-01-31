package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/inyourtime/noti-me-server/config"
	v1 "github.com/inyourtime/noti-me-server/internal/adapter/http/router/v1"
	"github.com/inyourtime/noti-me-server/internal/core/port"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type httpServer struct {
	cfg *config.Config
	app *echo.Echo
	db  *gorm.DB
}

func NewHttpServer(cfg *config.Config, db *gorm.DB) port.HttpServer {
	return &httpServer{
		cfg: cfg,
		app: echo.New(),
		db:  db,
	}
}

func (s *httpServer) Start() error {
	app := s.app

	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())

	apiV1 := app.Group("/v1")
	v1.NewRouter(apiV1, s.db)

	go s.listen()
	if err := s.gracefulShutdown(); err != nil {
		return err
	}

	return nil
}

func (s *httpServer) listen() {
	if err := s.app.Start(":" + s.cfg.Port); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatal(err)
	}
}

func (s *httpServer) gracefulShutdown() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.app.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
