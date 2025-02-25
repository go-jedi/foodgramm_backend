package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	"github.com/go-jedi/foodgrammm-backend/internal/middleware"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	userService service.UserService
	cookie      config.CookieConfig
	middleware  *middleware.Middleware
	logger      *logger.Logger
	validator   *validator.Validator
}

func NewHandler(
	userService service.UserService,
	cookie config.CookieConfig,
	middleware *middleware.Middleware,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		userService: userService,
		cookie:      cookie,
		middleware:  middleware,
		logger:      logger,
		validator:   validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/user")
	{
		api.POST("", h.create)
		api.POST("/list", h.middleware.Auth.AuthMiddleware, h.list)
		api.POST("/exists", h.exists)
		api.GET("/exists/id/:userID", h.existsByID)
		api.GET("/exists/telegram/:telegramID", h.existsByTelegramID)
		api.GET("/all", h.middleware.Auth.AuthMiddleware, h.all)
		api.GET("/count", h.middleware.Auth.AuthMiddleware, h.getUserCount)
		api.GET("/id/:userID", h.middleware.Auth.AuthMiddleware, h.getByID)
		api.GET("/telegram/:telegramID", h.middleware.Auth.AuthMiddleware, h.getByTelegramID)
		api.PUT("", h.middleware.Auth.AuthMiddleware, h.update)
		api.DELETE("/id/:userID", h.middleware.Auth.AuthMiddleware, h.deleteByID)
		api.DELETE("/telegram/:telegramID", h.middleware.Auth.AuthMiddleware, h.deleteByTelegramID)
	}
}
