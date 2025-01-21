package app

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/pkg/httpserver"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type App struct {
	cfg       config.Config
	logger    *logger.Logger
	validator *validator.Validator
	hs        *httpserver.HTTPServer
	db        *postgres.Postgres
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run application.
func (a *App) Run() error {
	return a.runHTTPServer()
}

// initDeps initialize deps.
func (a *App) initDeps(ctx context.Context) error {
	i := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initValidator,
		a.initPostgres,
		a.initHTTPServer,
		a.initModules,
	}

	for _, f := range i {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

// initConfig initialize config.
func (a *App) initConfig(_ context.Context) (err error) {
	a.cfg, err = config.GetConfig()
	if err != nil {
		return err
	}

	return nil
}

// initLogger initialize logger.
func (a *App) initLogger(_ context.Context) error {
	a.logger = logger.NewLogger(a.cfg.Logger)
	return nil
}

// initValidator initialize validator.
func (a *App) initValidator(_ context.Context) error {
	a.validator = validator.NewValidator()
	return nil
}

// initPostgres initialize postgres.
func (a *App) initPostgres(ctx context.Context) (err error) {
	a.db, err = postgres.NewPostgres(ctx, a.cfg.Postgres)
	if err != nil {
		return err
	}

	return nil
}

// initHTTPServer initialize http server.
func (a *App) initHTTPServer(_ context.Context) (err error) {
	a.hs, err = httpserver.NewHTTPServer(a.cfg.HTTPServer)
	if err != nil {
		return err
	}

	return nil
}

// initModules initialize modules.
func (a *App) initModules(ctx context.Context) error {
	// user.NewUser(a.hs.Engine, a.logger, a.validator, a.db, a.cache).Init(ctx)
	// captcha.NewCaptcha(a.hs.Engine, a.logger, a.validator).Init(ctx)

	return nil
}

// runHTTPServer run http server.
func (a *App) runHTTPServer() error {
	return a.hs.Start()
}
