package userblacklist

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	userblacklist "github.com/go-jedi/foodgramm_backend/internal/domain/user_blacklist"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Unblock a user
// @Description Unblock a user with the provided Telegram ID. Accessible only to administrators.
// @Tags UserBlacklist
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the admin user"
// @Success 200 {string} string "User unblocked successfully"
// @Failure 400 {object} userblacklist.ErrorResponse
// @Failure 401 {object} userblacklist.ErrorResponse
// @Failure 403 {object} userblacklist.ErrorResponse
// @Failure 409 {object} userblacklist.ErrorResponse
// @Failure 500 {object} userblacklist.ErrorResponse
// @Router /v1/user/blacklist/unblock/{telegramID} [delete]
func (h *Handler) unblockUser(c *gin.Context) {
	h.logger.Debug("[unblock user] execute handler")

	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, userblacklist.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.userBlackListService.UnblockUser(c.Request.Context(), telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserInBlackListDoesNotExist) {
			h.logger.Error("user in blacklist does not exist", "error", err)
			c.JSON(http.StatusConflict, userblacklist.ErrorResponse{
				Error:  "user in blacklist does not exist",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to unblock user", "error", err)
		c.JSON(http.StatusInternalServerError, userblacklist.ErrorResponse{
			Error:  "failed to unblock user",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
