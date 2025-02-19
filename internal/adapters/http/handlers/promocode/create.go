package promocode

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Create a new promo code
// @Description Creates a new promo code with the provided details.
// @Tags Promo code
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body promocode.CreateDTO true "Promo code creation data"
// @Success 201 {object} promocode.PromoCode "Created promo code details"
// @Failure 400 {object} promocode.ErrorResponse "Bad request due to invalid input or validation failure"
// @Failure 500 {object} promocode.ErrorResponse "Internal server error"
// @Router /v1/promo_code [post]
func (h *Handler) create(c *gin.Context) {
	h.logger.Debug("[create promo code] execute handler")

	var dto promocode.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, promocode.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, promocode.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.promoCodeService.Create(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrPromoCodeAlreadyExists) {
			c.JSON(http.StatusNotFound, promocode.ErrorResponse{
				Error:  "promo code already exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to create promo code", "error", err)
		c.JSON(http.StatusInternalServerError, promocode.ErrorResponse{
			Error:  "failed to create promo code",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
