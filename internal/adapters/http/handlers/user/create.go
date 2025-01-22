package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (h *Handler) Create(c *gin.Context) {
	var dto user.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request",
			"detail": err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request",
			"detail": err.Error(),
		})
		return
	}

	result, err := h.userService.Create(c, dto)
	if err != nil {
		h.logger.Error("failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to create user",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
