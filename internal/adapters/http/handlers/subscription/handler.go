package subscription

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/config"
	"github.com/go-jedi/foodgramm_backend/internal/middleware"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/validator"
)

type Handler struct {
	subscriptionService service.SubscriptionService
	cookie              config.CookieConfig
	middleware          *middleware.Middleware
	logger              *logger.Logger
	validator           *validator.Validator
}

func NewHandler(
	subscriptionService service.SubscriptionService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		subscriptionService: subscriptionService,
		cookie:              cookie,
		middleware:          middleware,
		logger:              logger,
		validator:           validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/subscription", h.middleware.Auth.AuthMiddleware)
	{
		api.GET("/telegram/:telegramID", h.getByTelegramID)
		api.GET("/exists/telegram/:telegramID", h.existsByTelegramID)
	}
}
