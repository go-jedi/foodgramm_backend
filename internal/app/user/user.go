package user

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	userRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/user"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	userService "github.com/go-jedi/foodgrammm-backend/internal/service/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type User struct {
	engine    *gin.Engine
	logger    *logger.Logger
	validator *validator.Validator
	db        *postgres.Postgres

	// user
	userRepository repository.UserRepository
	userService    service.UserService
	userHandler    *user.Handler
}

func NewUser(
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
	db *postgres.Postgres,
) *User {
	return &User{
		engine:    engine,
		logger:    logger,
		validator: validator,
		db:        db,
	}
}

func (u *User) Init(ctx context.Context) {
	_ = u.UserHandler(ctx)
}

func (u *User) UserRepository(_ context.Context) repository.UserRepository {
	if u.userRepository == nil {
		u.userRepository = userRepository.NewRepository(
			u.logger,
			u.db,
		)
	}

	return u.userRepository
}

func (u *User) UserService(ctx context.Context) service.UserService {
	if u.userService == nil {
		u.userService = userService.NewService(
			u.UserRepository(ctx),
			u.logger,
		)
	}

	return u.userService
}

func (u *User) UserHandler(ctx context.Context) *user.Handler {
	if u.userHandler == nil {
		u.userHandler = user.NewHandler(
			u.UserService(ctx),
			u.engine,
			u.logger,
			u.validator,
		)
	}

	return u.userHandler
}
