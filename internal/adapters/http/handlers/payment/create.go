package payment

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/payment"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Create a payment link
// @Description Creates a payment link based on the provided Telegram ID and payment type.
// @Tags Payment
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param payment body payment.CreateDTO true "Payment creation details"
// @Success 200 {string} string "Payment link created successfully"
// @Failure 400 {object} payment.ErrorResponse
// @Failure 404 {object} payment.ErrorResponse
// @Failure 500 {object} payment.ErrorResponse
// @Router /v1/payment/link [post]
func (h *Handler) create(c *gin.Context) {
	h.logger.Debug("[create a payment link] execute handler")

	var dto payment.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, payment.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, payment.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.paymentService.Create(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, payment.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to create payment link by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, payment.ErrorResponse{
			Error:  "failed to create payment link by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
