package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type WebSocket struct {
	paymentService service.PaymentService
	cookie         config.CookieConfig
	middleware     *middleware.Middleware
	logger         *logger.Logger
	validator      *validator.Validator
}

func NewWebSocket(
	paymentService service.PaymentService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *WebSocket {
	ws := &WebSocket{
		paymentService: paymentService,
		cookie:         cookie,
		middleware:     middleware,
		logger:         logger,
		validator:      validator,
	}

	ws.initRoutes(engine)

	return ws
}

func (ws *WebSocket) initRoutes(engine *gin.Engine) {
	api := engine.Group("ws/v1/payment", ws.middleware.Auth.AuthMiddleware)
	{
		api.GET("/check/:telegramID", ws.check)
	}
}
