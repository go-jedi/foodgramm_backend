package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (h *Handler) exists(c *gin.Context) {
	var dto user.ExistsDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "failed to bind body",
			"detail": err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "failed to validate struct",
			"detail": err.Error(),
		})
		return
	}

	isExists, err := h.userService.Exists(c, dto.TelegramID, dto.Username)
	if err != nil {
		h.logger.Error("failed to exists user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to exists user",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, isExists)
}
