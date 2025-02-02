package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	recipeService service.RecipeService
	cookie        config.CookieConfig
	middleware    *middleware.Middleware
	logger        *logger.Logger
	validator     *validator.Validator
}

func NewHandler(
	recipeService service.RecipeService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		recipeService: recipeService,
		cookie:        cookie,
		middleware:    middleware,
		logger:        logger,
		validator:     validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/recipe", h.middleware.Auth.AuthMiddleware)
	{
		api.GET("/free/telegram/:telegramID", h.getFreeRecipesByTelegramID)
	}
}
