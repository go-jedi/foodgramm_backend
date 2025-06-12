package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/internal/domain/auth"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// Refresh
// @Summary Refresh user token
// @Description Refresh the access token using the provided Telegram ID and refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body auth.RefreshDTO true "Refresh request body"
// @Success 200 {object} auth.RefreshResponse "Successful response with new tokens"
// @Failure 400 {object} auth.ErrorResponse "Bad request error"
// @Failure 500 {object} auth.ErrorResponse "Internal server error"
// @Router /v1/auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	h.logger.Debug("[refresh user token] execute handler")

	var dto auth.RefreshDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, auth.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, auth.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.authService.Refresh(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, auth.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to refresh token", "error", err)
		c.JSON(http.StatusInternalServerError, auth.ErrorResponse{
			Error:  "failed to refresh token",
			Detail: err.Error(),
		})
		return
	}

	c.SetCookie(
		h.cookie.Refresh.Name,
		result.RefreshToken,
		h.cookie.Refresh.MaxAge,
		h.cookie.Refresh.Path,
		h.cookie.Refresh.Domain,
		h.cookie.Refresh.Secure,
		h.cookie.Refresh.HTTPOnly,
	)

	c.JSON(http.StatusOK, result)
}
