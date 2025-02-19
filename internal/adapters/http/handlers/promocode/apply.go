package promocode

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Apply a promo code
// @Description Applies a promo code for a user with the provided details.
// @Tags Promo code
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body promocode.ApplyDTO true "Promo code application data"
// @Success 201 {object} promocode.ApplyResponse "Successfully applied promo code details"
// @Failure 400 {object} promocode.ErrorResponse "Bad request due to invalid input or validation failure"
// @Failure 404 {object} promocode.ErrorResponse "User does not exist or promo code is not valid for the user"
// @Failure 500 {object} promocode.ErrorResponse "Internal server error"
// @Router /v1/promo_code/apply [post]
func (h *Handler) apply(c *gin.Context) {
	h.logger.Debug("[apply promo code] execute handler")

	var dto promocode.ApplyDTO
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

	result, err := h.promoCodeService.Apply(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, promocode.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}
		if errors.Is(err, apperrors.ErrPromoCodeIsNotValidForUser) {
			c.JSON(http.StatusNotFound, promocode.ErrorResponse{
				Error:  "promo code is not valid for user",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to apply promo code", "error", err)
		c.JSON(http.StatusInternalServerError, promocode.ErrorResponse{
			Error:  "failed to apply promo code",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
