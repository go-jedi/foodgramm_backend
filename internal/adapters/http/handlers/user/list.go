package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (h *Handler) list(c *gin.Context) {
	var dto user.ListDTO
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

	result, err := h.userService.List(c, dto)
	if err != nil {
		h.logger.Error("failed to get list users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to get list users",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
