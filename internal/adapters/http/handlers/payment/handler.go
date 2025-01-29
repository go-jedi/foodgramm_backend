package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type Handler struct {
	paymentService service.PaymentService
	logger         *logger.Logger
	validator      *validator.Validator
}

func NewHandler(
	paymentService service.PaymentService,
	engine *gin.Engine,
	logger *logger.Logger,
	validator *validator.Validator,
) *Handler {
	h := &Handler{
		paymentService: paymentService,
		logger:         logger,
		validator:      validator,
	}

	h.initRoutes(engine)

	return h
}

func (h *Handler) initRoutes(engine *gin.Engine) {
	api := engine.Group("/v1/payment")
	{
		api.POST("/create/link", h.create)
		api.GET("/check/status", h.checkStatus)
	}
}
