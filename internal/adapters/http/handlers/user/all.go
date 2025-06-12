package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
)

// @Summary Get all users
// @Description Retrieve a list of all users.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Success 200 {array} user.User
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/all [get]
func (h *Handler) all(c *gin.Context) {
	h.logger.Debug("[get all users] execute handler")

	result, err := h.userService.All(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to get all users", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to get all users",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
