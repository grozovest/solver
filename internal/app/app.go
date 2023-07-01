package app

import (
	"context"
	"log"
	"net/http"

	"solver/internal/config"
	"solver/internal/handler"
	"solver/internal/migrate"
	"solver/internal/router"
	"solver/internal/storage"
	"solver/internal/usecase"
	"solver/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type App struct {
	HTTPServer *http.Server
	logger     *logger.Logger
}

func New(cfg config.Config) (*App, error) {
	ctx := context.Background()

	lg, err := logger.New(cfg.Debug)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	err = migrate.Migrate(cfg.Database, migrate.Migrations)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(ctx, cfg.Database)
	if err != nil {
		lg.Logger.Error(err.Error())

		return nil, err
	}

	repo := storage.New(pool, lg)
	useCase := usecase.New(lg, repo)
	h := handler.New(lg, useCase)

	srv := &http.Server{
		Handler: router.New(h),
		Addr:    cfg.Address,
	}

	return &App{
		HTTPServer: srv,
		logger:     lg,
	}, nil
}

func (app *App) Run() error {
	app.logger.Logger.Info("server started", zap.String("addr", app.HTTPServer.Addr))
	return app.HTTPServer.ListenAndServe()
}
