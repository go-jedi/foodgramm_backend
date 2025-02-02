package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	authService service.AuthService
	cookie      config.CookieConfig
	middleware  *middleware.Middleware
	logger      *logger.Logger
	validator   *validator.Validator
}

func NewHandler(
	authService service.AuthService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		authService: authService,
		cookie:      cookie,
		middleware:  middleware,
		logger:      logger,
		validator:   validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/auth")
	{
		api.POST("/signin", h.SignIn)
		api.POST("/check", h.middleware.Auth.AuthMiddleware, h.Check)
		api.POST("/refresh", h.middleware.Auth.AuthMiddleware, h.Refresh)
	}
}
