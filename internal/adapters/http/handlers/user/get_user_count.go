package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (h *Handler) getUserCount(c *gin.Context) {
	h.logger.Debug("[get user count] execute handler")

	result, err := h.userService.GetUserCount(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to get user count", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to get user count",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
