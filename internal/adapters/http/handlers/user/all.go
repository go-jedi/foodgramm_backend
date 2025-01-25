package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) all(c *gin.Context) {
	result, err := h.userService.All(c)
	if err != nil {
		h.logger.Error("failed to get all users", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to get all users",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
