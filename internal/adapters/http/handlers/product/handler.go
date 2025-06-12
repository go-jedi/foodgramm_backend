package product

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/config"
	"github.com/go-jedi/foodgramm_backend/internal/middleware"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/validator"
)

type Handler struct {
	productService service.ProductService
	cookie         config.CookieConfig
	middleware     *middleware.Middleware
	logger         *logger.Logger
	validator      *validator.Validator
}

func NewHandler(
	productService service.ProductService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		productService: productService,
		cookie:         cookie,
		middleware:     middleware,
		logger:         logger,
		validator:      validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/product", h.middleware.Auth.AuthMiddleware)
	{
		api.POST("/exclude/telegram/id", h.addExcludeProductsByTelegramID)
		api.GET("/exclude/telegram/:telegramID", h.getExcludeProductsByTelegramID)
		api.DELETE("/exclude/telegram/:telegramID", h.deleteExcludeProductsByTelegramID)
	}
}
