package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (h *Handler) getByID(c *gin.Context) {
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

	result, err := h.userService.GetByID(c, userIDInt)
	if err != nil {
		h.logger.Error("failed to get user by id", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to get user by id",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
