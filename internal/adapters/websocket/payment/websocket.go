package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type WebSocketHandler struct {
	subscriptionService service.SubscriptionService
	cookie              config.CookieConfig
	websocket           config.WebSocketConfig
	middleware          *middleware.Middleware
	logger              *logger.Logger
	validator           *validator.Validator
}

func NewWebSocketHandler(
	subscriptionService service.SubscriptionService,
	cookie config.CookieConfig,
	websocket config.WebSocketConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *WebSocketHandler {
	h := &WebSocketHandler{
		subscriptionService: subscriptionService,
		cookie:              cookie,
		websocket:           websocket,
		middleware:          middleware,
		logger:              logger,
		validator:           validator,
	}

	h.initRoutes(engine)

	return h
}

func (wsh *WebSocketHandler) initRoutes(engine *gin.Engine) {
	api := engine.Group("ws/v1/payment")
	{
		api.GET("/check/:telegramID", wsh.check)
	}
}
