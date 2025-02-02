package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Check if a user exists by Telegram ID
// @Description Check if a user exists by their unique Telegram ID.
// @Tags User
// @Accept json
// @Produce json
// @Param telegramID path string true "Telegram ID"
// @Success 200 {boolean} boolean "Boolean flag indicating if the user exists"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/exists/telegram/{telegramID} [get]
func (h *Handler) existsByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.userService.ExistsByTelegramID(c, telegramID)
	if err != nil {
		h.logger.Error("failed to exists user by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to exists user by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
