package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/product"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Get Exclude Products by Telegram ID
// @Description Get excluded products for a user by their unique Telegram ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID"
// @Success 200 {object} product.UserExcludedProducts "Excluded products for the user"
// @Failure 400 {object} product.ErrorResponse
// @Failure 500 {object} product.ErrorResponse
// @Router /v1/product/exclude/telegram/{telegramID} [get]
func (h *Handler) getExcludeProductsByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, product.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.productService.GetExcludeProductsByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		h.logger.Error("failed to get exclude products by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, product.ErrorResponse{
			Error:  "failed to get exclude products by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
