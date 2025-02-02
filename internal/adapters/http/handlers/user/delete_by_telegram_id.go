package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Delete a user by Telegram ID
// @Description Delete a user by their unique Telegram ID.
// @Tags User
// @Accept json
// @Produce json
// @Param telegramID path string true "Telegram ID"
// @Success 200 {string} string "Deleted Telegram ID"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/telegram/{telegramID} [delete]
func (h *Handler) deleteByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.userService.DeleteByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		h.logger.Error("failed to delete user by telegramID", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to delete user by telegramID",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
