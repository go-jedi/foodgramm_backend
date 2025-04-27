package admin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/admin"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Add a new admin user
// @Description Add a new admin user with the provided telegram ID. Accessible only to administrators.
// @Tags Admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the admin user"
// @Success 200 {object} Admin
// @Failure 400 {object} admin.ErrorResponse
// @Failure 401 {object} admin.ErrorResponse
// @Failure 403 {object} admin.ErrorResponse
// @Failure 409 {object} admin.ErrorResponse
// @Failure 500 {object} admin.ErrorResponse
// @Router /v1/admin/add/{telegramID} [get]
func (h *Handler) addAdminUser(c *gin.Context) {
	h.logger.Debug("[add a new admin user] execute handler")

	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, admin.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.adminService.AddAdminUser(c.Request.Context(), telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrAdminAlreadyExists) {
			h.logger.Error("admin already exists", "error", err)
			c.JSON(http.StatusConflict, admin.ErrorResponse{
				Error:  "admin already exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to add admin user", "error", err)
		c.JSON(http.StatusInternalServerError, admin.ErrorResponse{
			Error:  "failed to add admin user",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
