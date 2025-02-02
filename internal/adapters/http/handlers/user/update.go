package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

// @Summary Update a user
// @Description Update a user's information.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body user.UpdateDTO true "User update data"
// @Success 200 {object} user.User "Updated user details"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user [put]
func (h *Handler) update(c *gin.Context) {
	var dto user.UpdateDTO
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

	result, err := h.userService.Update(c.Request.Context(), dto)
	if err != nil {
		h.logger.Error("failed to update user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to update user",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
