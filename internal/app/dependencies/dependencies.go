package dependencies

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/product"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Dependencies struct {
	engine    *gin.Engine
	logger    *logger.Logger
	validator *validator.Validator
	db        *postgres.Postgres
	cache     *redis.Redis

	// user
	userRepository repository.UserRepository
	userService    service.UserService
	userHandler    *user.Handler

	// product
	productRepository repository.ProductRepository
	productService    service.ProductService
	productHandler    *product.Handler
}

func NewDependencies(
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
	db *postgres.Postgres,
	cache *redis.Redis,
) *Dependencies {
	d := &Dependencies{
		engine:    engine,
		logger:    logger,
		validator: validator,
		db:        db,
		cache:     cache,
	}

	d.Init()

	return d
}

func (d *Dependencies) Init() {
	_ = d.UserHandler()
	_ = d.ProductHandler()
}
