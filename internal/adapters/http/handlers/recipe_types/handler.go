package recipetypes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/config"
	"github.com/go-jedi/foodgramm_backend/internal/middleware"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/validator"
)

type Handler struct {
	recipeTypesService service.RecipeTypesService
	cookie             config.CookieConfig
	middleware         *middleware.Middleware
	logger             *logger.Logger
	validator          *validator.Validator
}

func NewHandler(
	recipeTypesService service.RecipeTypesService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		recipeTypesService: recipeTypesService,
		cookie:             cookie,
		middleware:         middleware,
		logger:             logger,
		validator:          validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/recipe_types", h.middleware.Auth.AuthMiddleware)
	{
		api.POST("", h.create)
		api.GET("/all", h.all)
		api.GET("/id/:recipeTypeID", h.getByID)
		api.PATCH("", h.update)
		api.DELETE("/id/:recipeTypeID", h.deleteByID)
	}
}
