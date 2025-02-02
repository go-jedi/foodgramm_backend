package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	authService service.AuthService
	logger      *logger.Logger
	validator   *validator.Validator
}

func NewHandler(
	authService service.AuthService,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		authService: authService,
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
		api.POST("/check", h.Check)
		api.POST("/refresh", h.Refresh)
	}
}
