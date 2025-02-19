package promocode

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	promoCodeService service.PromoCodeService
	cookie           config.CookieConfig
	middleware       *middleware.Middleware
	logger           *logger.Logger
	validator        *validator.Validator
}

func NewHandler(
	promoCodeService service.PromoCodeService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		promoCodeService: promoCodeService,
		cookie:           cookie,
		middleware:       middleware,
		logger:           logger,
		validator:        validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/promo_code", h.middleware.Auth.AuthMiddleware)
	{
		api.POST("", h.create)
		api.POST("/validate", h.isPromoCodeValidForUser)
		api.POST("/apply", h.apply)
	}
}
