package userblacklist

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userblacklist "github.com/go-jedi/foodgrammm-backend/internal/domain/user_blacklist"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Check if a user exists in the blacklist by Telegram ID
// @Description Check if a user exists in the blacklist with the provided Telegram ID. Accessible only to administrators.
// @Tags UserBlacklist
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the user to check"
// @Success 200 {boolean} bool "Indicates if the user exists in the blacklist"
// @Failure 400 {object} userblacklist.ErrorResponse
// @Failure 401 {object} userblacklist.ErrorResponse
// @Failure 403 {object} userblacklist.ErrorResponse
// @Failure 500 {object} userblacklist.ErrorResponse
// @Router /v1/user/blacklist/exists/{telegramID} [get]
func (h *Handler) exists(c *gin.Context) {
	h.logger.Debug("[check user exists in blacklist by telegram id] execute handler")

	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, userblacklist.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.userBlackListService.Exists(c.Request.Context(), telegramID)
	if err != nil {
		h.logger.Error("failed to exists in blacklist by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, userblacklist.ErrorResponse{
			Error:  "failed to exists in blacklist by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
