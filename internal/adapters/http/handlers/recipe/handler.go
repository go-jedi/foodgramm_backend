package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/config"
	"github.com/go-jedi/foodgramm_backend/internal/middleware"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/validator"
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
		api.POST("/generate", h.generateRecipe)
		api.POST("/list", h.getListRecipesByTelegramID)
		api.GET("/telegram/:telegramID", h.getRecipesByTelegramID)
		api.GET("/free/telegram/:telegramID", h.getFreeRecipesByTelegramID)
	}
}
