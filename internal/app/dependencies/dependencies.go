package dependencies

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	recipeofdayscron "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/cron/recipe_of_days"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/auth"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/payment"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/product"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe"
	recipeofdayshandler "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe_of_days"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user"
	"github.com/go-jedi/foodgrammm-backend/internal/client"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/internal/templates"
	"github.com/go-jedi/foodgrammm-backend/pkg/jwt"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Dependencies struct {
	cookie     config.CookieConfig
	worker     config.WorkerConfig
	engine     *gin.Engine
	middleware *middleware.Middleware
	logger     *logger.Logger
	validator  *validator.Validator
	jwt        *jwt.JWT
	templates  *templates.Templates
	parser     *parser.Parser
	client     *client.Client
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

	// recipe of days
	recipeOfDaysRepository repository.RecipeOfDaysRepository
	recipeOfDaysService    service.RecipeOfDaysService
	recipeOfDaysHandler    *recipeofdayshandler.Handler

	// payment
	paymentService service.PaymentService
	paymentHandler *payment.Handler

	//	cron recipe of days
	recipeOfDaysCron *recipeofdayscron.Cron
}

func NewDependencies(
	ctx context.Context,
	cookie config.CookieConfig,
	worker config.WorkerConfig,
	engine *gin.Engine,
	middleware *middleware.Middleware,
	logger *logger.Logger,
	validator *validator.Validator,
	jwt *jwt.JWT,
	templates *templates.Templates,
	parser *parser.Parser,
	client *client.Client,
	db *postgres.Postgres,
	cache *redis.Redis,
) *Dependencies {
	d := &Dependencies{
		cookie:     cookie,
		worker:     worker,
		engine:     engine,
		middleware: middleware,
		logger:     logger,
		validator:  validator,
		jwt:        jwt,
		templates:  templates,
		parser:     parser,
		client:     client,
		db:         db,
		cache:      cache,
	}

	d.initHandler()
	d.initCron(ctx)

	return d
}

func (d *Dependencies) initHandler() {
	_ = d.AuthHandler()
	_ = d.UserHandler()
	_ = d.ProductHandler()
	_ = d.RecipeHandler()
	_ = d.SubscriptionHandler()
	_ = d.RecipeOfDaysHandler()
	_ = d.PaymentHandler()
}

func (d *Dependencies) initCron(ctx context.Context) {
	_ = d.RecipeOfDaysCron(ctx)
}
