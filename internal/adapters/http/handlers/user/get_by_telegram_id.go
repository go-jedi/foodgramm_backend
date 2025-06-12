package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Get a user by Telegram ID
// @Description Retrieve a user by their unique Telegram ID.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID"
// @Success 200 {object} user.User "User details"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/telegram/{telegramID} [get]
func (h *Handler) getByTelegramID(c *gin.Context) {
	h.logger.Debug("[get user by telegram id] execute handler")

	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.userService.GetByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		h.logger.Error("failed to get user by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to get user by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
