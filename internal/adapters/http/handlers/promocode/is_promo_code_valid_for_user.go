package promocode

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/internal/domain/promocode"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Check if a promo code is valid for a user
// @Description Checks if a promo code is valid for the specified user.
// @Tags Promo code
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body promocode.IsPromoCodeValidForUserDTO true "Promo code validation data"
// @Success 200 {boolean} boolean "Returns true if the promo code is valid for the user, otherwise false"
// @Failure 400 {object} promocode.ErrorResponse "Bad request due to invalid input or validation failure"
// @Failure 404 {object} promocode.ErrorResponse "User does not exist"
// @Failure 500 {object} promocode.ErrorResponse "Internal server error"
// @Router /v1/promo_code/validate [post]
func (h *Handler) isPromoCodeValidForUser(c *gin.Context) {
	h.logger.Debug("[is promo code valid for user] execute handler")

	var dto promocode.IsPromoCodeValidForUserDTO
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

	result, err := h.promoCodeService.IsPromoCodeValidForUser(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, promocode.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to check valid promo code for user", "error", err)
		c.JSON(http.StatusInternalServerError, promocode.ErrorResponse{
			Error:  "failed to check valid promo code for user",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
