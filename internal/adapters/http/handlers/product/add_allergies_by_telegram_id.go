package product

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Add Allergies by Telegram ID
// @Description Add or update user allergies based on Telegram ID
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param product body product.AddAllergiesByTelegramIDDTO true "Add Allergies Request"
// @Success 200 {object} product.UserExcludedProducts "User Excluded Products Response"
// @Failure 400 {object} product.ErrorResponse "Bad Request"
// @Failure 404 {object} product.ErrorResponse "User Not Found"
// @Failure 500 {object} product.ErrorResponse "Internal Server Error"
// @Router /v1/product/allergy [post]
func (h *Handler) addAllergiesByTelegramID(c *gin.Context) {
	var dto product.AddAllergiesByTelegramIDDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.productService.AddAllergiesByTelegramID(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, product.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to add allergies by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, product.ErrorResponse{
			Error:  "failed to add allergies by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
