package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (h *Handler) getExcludeProductsByUserID(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		h.logger.Error("failed to get param userID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "failed to get param userID",
			"detail": apperrors.ErrParamIsRequired,
		})
		return
	}

	const (
		base    = 10
		bitSize = 64
	)

	userIDInt, err := strconv.ParseInt(userID, base, bitSize)
	if err != nil {
		h.logger.Error("failed parse string to int64", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "failed parse string to int64",
			"detail": err.Error(),
		})
		return
	}

	if userIDInt <= 0 {
		h.logger.Error("userID is less than or equal to zero", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "userID must be greater than zero",
			"detail": "the provided userID is less than or equal to zero, which is not allowed",
		})
		return
	}

	result, err := h.productService.GetExcludeProductsByUserID(c, userIDInt)
	if err != nil {
		h.logger.Error("failed to get exclude products by user id", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to get exclude products by user id",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
