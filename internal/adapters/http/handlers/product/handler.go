package product

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	productService service.ProductService
	logger         *logger.Logger
	validator      *validator.Validator
}

func NewHandler(
	productService service.ProductService,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		productService: productService,
		logger:         logger,
		validator:      validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/product")
	{
		api.POST("/exclude/id", h.excludeProductsByID)
		api.POST("/exclude/telegram", h.excludeProductsByTelegramID)
	}
}
