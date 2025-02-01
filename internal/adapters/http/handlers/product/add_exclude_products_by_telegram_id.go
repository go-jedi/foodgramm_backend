package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
)

// @Summary Add Exclude Products by Telegram ID
// @Description Add excluded products for a user by their unique Telegram ID.
// @Tags users
// @Accept json
// @Produce json
// @Param request body product.AddExcludeProductsByTelegramIDDTO true "Exclude products data"
// @Success 200 {object} product.UserExcludedProducts "Excluded products for the user"
// @Failure 400 {object} product.ErrorResponse
// @Failure 500 {object} product.ErrorResponse
// @Router /v1/product/exclude/telegram/id [post]
func (h *Handler) addExcludeProductsByTelegramID(c *gin.Context) {
	var dto product.AddExcludeProductsByTelegramIDDTO
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

	result, err := h.productService.AddExcludeProductsByTelegramID(c, dto)
	if err != nil {
		h.logger.Error("failed to add exclude products by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, product.ErrorResponse{
			Error:  "failed to add exclude products by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
