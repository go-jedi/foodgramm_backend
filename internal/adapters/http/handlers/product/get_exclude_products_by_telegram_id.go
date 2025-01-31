package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (h *Handler) getExcludeProductsByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "failed to get param telegramID",
			"detail": apperrors.ErrParamIsRequired,
		})
		return
	}

	result, err := h.productService.GetExcludeProductsByTelegramID(c, telegramID)
	if err != nil {
		h.logger.Error("failed to get exclude products by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to get exclude products by telegram id",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
