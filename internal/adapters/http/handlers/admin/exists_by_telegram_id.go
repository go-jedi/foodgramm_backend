package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/admin"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Check if an admin exists by Telegram ID
// @Description Check if an admin exists with the provided Telegram ID. Accessible only to administrators.
// @Tags Admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the user to check"
// @Success 200 {boolean} boolean
// @Failure 400 {object} admin.ErrorResponse
// @Failure 401 {object} admin.ErrorResponse
// @Failure 403 {object} admin.ErrorResponse
// @Failure 500 {object} admin.ErrorResponse
// @Router /v1/admin/exists/{telegramID} [get]
func (h *Handler) existsByTelegramID(c *gin.Context) {
	h.logger.Debug("[check admin exists by telegram id] execute handler")

	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, admin.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.adminService.ExistsByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		h.logger.Error("failed to exists admin by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, admin.ErrorResponse{
			Error:  "failed to exists admin by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
