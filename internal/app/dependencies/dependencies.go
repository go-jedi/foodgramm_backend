package dependencies

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	recipeofdayscron "github.com/go-jedi/foodgrammm-backend/internal/adapters/cron/recipe_of_days"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/admin"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/auth"
	clientassets "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/client_assets"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/payment"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/product"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/promocode"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe_event"
	recipeofdayshandler "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe_of_days"
	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/recipe_types"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/subscription"
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user"
	userblacklist "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user_blacklist"
	paymentwebsocket "github.com/go-jedi/foodgrammm-backend/internal/adapters/websocket/payment"
	"github.com/go-jedi/foodgrammm-backend/internal/client"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/internal/templates"
	fileserver "github.com/go-jedi/foodgrammm-backend/pkg/file_server"
	"github.com/go-jedi/foodgrammm-backend/pkg/jwt"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Dependencies struct {
	cookie     config.CookieConfig
	websocket  config.WebSocketConfig
	worker     config.WorkerConfig
	engine     *gin.Engine
	fileServer *fileserver.FileServer
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

	// recipe types
	recipeTypesRepository repository.RecipeTypesRepository
	recipeTypesService    service.RecipeTypesService
	recipeTypesHandler    *recipetypes.Handler

	// recipe event
	recipeEventRepository repository.RecipeEventRepository
	recipeEventService    service.RecipeEventService
	recipeEventHandler    *recipeevent.Handler

	// payment
	paymentService service.PaymentService
	paymentHandler *payment.Handler

	// promo code
	promoCodeRepository repository.PromoCodeRepository
	promoCodeService    service.PromoCodeService
	promoCodeHandler    *promocode.Handler

	// client assets
	clientAssetsRepository repository.ClientAssetsRepository
	clientAssetsService    service.ClientAssetsService
	clientAssetsHandler    *clientassets.Handler

	// admin
	adminRepository repository.AdminRepository
	adminService    service.AdminService
	adminHandler    *admin.Handler

	// user blackList
	userBlackListRepository repository.UserBlackListRepository
	userBlackListService    service.UserBlackListService
	userBlackListHandler    *userblacklist.Handler

	// websocket
	paymentWebSocketHandler *paymentwebsocket.WebSocketHandler

	//	cron
	recipeOfDaysCron *recipeofdayscron.Cron
}

func NewDependencies(
	ctx context.Context,
	cookie config.CookieConfig,
	websocket config.WebSocketConfig,
	worker config.WorkerConfig,
	engine *gin.Engine,
	fileServer *fileserver.FileServer,
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
		websocket:  websocket,
		worker:     worker,
		engine:     engine,
		fileServer: fileServer,
		logger:     logger,
		validator:  validator,
		jwt:        jwt,
		templates:  templates,
		parser:     parser,
		client:     client,
		db:         db,
		cache:      cache,
	}

	d.initMiddleware()
	d.initHandler()
	d.initWebSocket()
	d.initCron(ctx)

	return d
}

// initMiddleware initialize middlewares.
func (d *Dependencies) initMiddleware() {
	d.middleware = middleware.New(
		d.AdminService(),
		d.jwt,
	)
}

// initHandler initialize handlers.
func (d *Dependencies) initHandler() {
	_ = d.AuthHandler()
	_ = d.UserHandler()
	_ = d.ProductHandler()
	_ = d.RecipeHandler()
	_ = d.SubscriptionHandler()
	_ = d.RecipeOfDaysHandler()
	_ = d.RecipeTypesHandler()
	_ = d.RecipeEventHandler()
	_ = d.PaymentHandler()
	_ = d.PromoCodeHandler()
	_ = d.ClientAssetsHandler()
	_ = d.AdminHandler()
	_ = d.UserBlackListHandler()
}

// initWebSocket initialize web sockets.
func (d *Dependencies) initWebSocket() {
	_ = d.PaymentWebSocket()
}

// initCron initialize crons.
func (d *Dependencies) initCron(ctx context.Context) {
	_ = d.RecipeOfDaysCron(ctx)
}
