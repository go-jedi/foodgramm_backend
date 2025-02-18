package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

// @Summary Check if a user exists
// @Description Check if a user exists by Telegram ID and Username.
// @Tags User
// @Accept json
// @Produce json
// @Param request body user.ExistsDTO true "User data"
// @Success 200 {boolean} boolean "Boolean flag indicating if the user exists"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/exists [post]
func (h *Handler) exists(c *gin.Context) {
	h.logger.Debug("[check a user exists] execute handler")

	var dto user.ExistsDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.userService.Exists(c.Request.Context(), dto.TelegramID, dto.Username)
	if err != nil {
		h.logger.Error("failed to exists user", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to exists user",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
