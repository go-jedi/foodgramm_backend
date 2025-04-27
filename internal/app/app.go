package app

import (
	"context"
	"net/http"

	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/app/dependencies"
	"github.com/go-jedi/foodgrammm-backend/internal/client"
	"github.com/go-jedi/foodgrammm-backend/internal/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/templates"
	"github.com/go-jedi/foodgrammm-backend/pkg/bcrypt"
	fileserver "github.com/go-jedi/foodgrammm-backend/pkg/file_server"
	"github.com/go-jedi/foodgrammm-backend/pkg/httpserver"
	"github.com/go-jedi/foodgrammm-backend/pkg/jwt"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
	"github.com/go-jedi/foodgrammm-backend/pkg/uid"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type App struct {
	cfg          config.Config
	logger       *logger.Logger
	validator    *validator.Validator
	bcrypt       *bcrypt.Bcrypt
	uid          *uid.UID
	jwt          *jwt.JWT
	templates    *templates.Templates
	parser       *parser.Parser
	client       *client.Client
	hs           *httpserver.HTTPServer
	fileServer   *fileserver.FileServer
	db           *postgres.Postgres
	cache        *redis.Redis
	dependencies *dependencies.Dependencies
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
		a.initBcrypt,
		a.initUID,
		a.initJWT,
		a.initTemplates,
		a.initParser,
		a.initClient,
		a.initPostgres,
		a.initRedis,
		a.initHTTPServer,
		a.initFileServer,
		a.initDependencies,
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

	return
}

// initLogger initialize logger.
func (a *App) initLogger(_ context.Context) error {
	a.logger = logger.New(a.cfg.Logger)
	return nil
}

// initValidator initialize validator.
func (a *App) initValidator(_ context.Context) error {
	a.validator = validator.New()
	return nil
}

// initBcrypt initialize bcrypt.
func (a *App) initBcrypt(_ context.Context) (err error) {
	a.bcrypt, err = bcrypt.NewBcryptWithCost(a.cfg.Bcrypt)
	if err != nil {
		return err
	}

	return
}

// initUID initialize uid.
func (a *App) initUID(_ context.Context) (err error) {
	a.uid, err = uid.New(uid.Option{
		Chars: a.cfg.UID.Chars,
		Count: a.cfg.UID.Count,
	})
	if err != nil {
		return err
	}

	return
}

// initJWT initialize jwt.
func (a *App) initJWT(_ context.Context) (err error) {
	a.jwt, err = jwt.New(a.cfg.JWT, a.uid)
	if err != nil {
		return err
	}

	return
}

// initTemplates initialize templates.
func (a *App) initTemplates(_ context.Context) error {
	a.templates = templates.NewTemplates()

	return nil
}

// initParser initialize parser.
func (a *App) initParser(_ context.Context) error {
	a.parser = parser.NewParser()

	return nil
}

// initClient initialize client.
func (a *App) initClient(_ context.Context) (err error) {
	a.client, err = client.NewClient(a.cfg.Client)
	if err != nil {
		return err
	}

	return
}

// initPostgres initialize postgres.
func (a *App) initPostgres(ctx context.Context) (err error) {
	a.db, err = postgres.New(ctx, a.cfg.Postgres)
	if err != nil {
		return err
	}

	return
}

// initRedis initialize redis.
func (a *App) initRedis(ctx context.Context) (err error) {
	a.cache, err = redis.New(ctx, a.cfg.Redis, a.db)
	if err != nil {
		return err
	}

	return
}

// initHTTPServer initialize http server.
func (a *App) initHTTPServer(_ context.Context) (err error) {
	a.hs, err = httpserver.New(a.cfg.HTTPServer)
	if err != nil {
		return err
	}

	return
}

// initFileServer initialize file server.
func (a *App) initFileServer(_ context.Context) error {
	a.fileServer = fileserver.New(a.cfg.FileServer)

	a.hs.Engine.StaticFS(a.cfg.FileServer.URL, http.Dir(a.cfg.FileServer.Dir))

	return nil
}

// initDependencies initialize dependencies.
func (a *App) initDependencies(ctx context.Context) error {
	a.dependencies = dependencies.NewDependencies(
		ctx,
		a.cfg.Cookie,
		a.cfg.WebSocket,
		a.cfg.Worker,
		a.hs.Engine,
		a.fileServer,
		a.logger,
		a.validator,
		a.jwt,
		a.templates,
		a.parser,
		a.client,
		a.db,
		a.cache,
	)

	return nil
}

// runHTTPServer run http server.
func (a *App) runHTTPServer() error {
	return a.hs.Start()
}
