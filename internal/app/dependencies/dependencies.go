package dependencies

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/auth"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/product"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/jwt"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Dependencies struct {
	cookie     config.CookieConfig
	engine     *gin.Engine
	middleware *middleware.Middleware
	logger     *logger.Logger
	validator  *validator.Validator
	jwt        *jwt.JWT
	db         *postgres.Postgres
	cache      *redis.Redis

	// auth
	authService service.AuthService
	authHandler *auth.Handler

	// user
	userRepository repository.UserRepository
	userService    service.UserService
	userHandler    *user.Handler

	// product
	productRepository repository.ProductRepository
	productService    service.ProductService
	productHandler    *product.Handler

	// recipe
	recipeRepository repository.RecipeRepository
	recipeService    service.RecipeService
	recipeHandler    *recipe.Handler

	// subscription
	subscriptionRepository repository.SubscriptionRepository
	subscriptionService    service.SubscriptionService
	subscriptionHandler    *subscription.Handler
}

func NewDependencies(
	cookie config.CookieConfig,
	engine *gin.Engine,
	middleware *middleware.Middleware,
	logger *logger.Logger,
	validator *validator.Validator,
	jwt *jwt.JWT,
	db *postgres.Postgres,
	cache *redis.Redis,
) *Dependencies {
	d := &Dependencies{
		cookie:     cookie,
		engine:     engine,
		middleware: middleware,
		logger:     logger,
		validator:  validator,
		jwt:        jwt,
		db:         db,
		cache:      cache,
	}

	d.Init()

	return d
}

func (d *Dependencies) Init() {
	_ = d.AuthHandler()
	_ = d.UserHandler()
	_ = d.ProductHandler()
	_ = d.RecipeHandler()
	_ = d.SubscriptionHandler()
}
