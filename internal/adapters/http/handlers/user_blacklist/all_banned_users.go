package userblacklist

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userblacklist "github.com/go-jedi/foodgramm_backend/internal/domain/user_blacklist"
)

// @Summary Get all banned users
// @Description Get a list of all banned users. Accessible only to administrators.
// @Tags UserBlacklist
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the admin user"
// @Success 200 {array} userblacklist.UsersBlackList "List of all banned users"
// @Failure 401 {object} userblacklist.ErrorResponse
// @Failure 403 {object} userblacklist.ErrorResponse
// @Failure 500 {object} userblacklist.ErrorResponse
// @Router /v1/user/blacklist/all [get]
func (h *Handler) allBannedUsers(c *gin.Context) {
	h.logger.Debug("[get all banned users] execute handler")

	result, err := h.userBlackListService.AllBannedUsers(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to get banned users", "error", err)
		c.JSON(http.StatusInternalServerError, userblacklist.ErrorResponse{
			Error:  "failed to get banned users",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
