package subscription

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Check if subscription exists by Telegram ID
// @Description Checks whether a subscription exists for a user with the specified Telegram ID.
// @Tags Subscription
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the user"
// @Success 200 {object} bool "Indicates whether the subscription exists"
// @Failure 400 {object} subscription.ErrorResponse "Invalid or missing 'telegramID' parameter"
// @Failure 404 {object} subscription.ErrorResponse "User not found"
// @Failure 500 {object} subscription.ErrorResponse "Internal server error"
// @Router /v1/subscription/exists/telegram/{telegramID} [get]
func (h *Handler) existsByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, subscription.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.subscriptionService.ExistsByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, subscription.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to exists subscription by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, subscription.ErrorResponse{
			Error:  "failed to exists subscription by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
