package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	userService service.UserService
	logger      *logger.Logger
	validator   *validator.Validator
}

func NewHandler(
	userService service.UserService,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		userService: userService,
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
		api.POST("/list", h.list)
		api.POST("/exists", h.exists)
		api.GET("/all", h.all)
		api.GET("/id/:userID", h.getByID)
		api.GET("/telegram/:telegramID", h.getByTelegramID)
	}
}
