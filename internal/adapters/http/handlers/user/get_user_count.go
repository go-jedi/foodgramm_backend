package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

// @Summary Get user count
// @Description Retrieve the total number of users.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Success 200 {integer} int64 "Returns the total number of users"
// @Failure 500 {object} user.ErrorResponse "Internal server error"
// @Router /v1/user/count [get]
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
