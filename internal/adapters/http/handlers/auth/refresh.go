package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
)

// @Summary Refresh user token
// @Description Refresh the access token using the provided Telegram ID and refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body auth.RefreshDTO true "Refresh request body"
// @Success 200 {object} auth.RefreshResponse "Successful response with new tokens"
// @Failure 400 {object} auth.ErrorResponse "Bad request error"
// @Failure 500 {object} auth.ErrorResponse "Internal server error"
// @Router /v1/auth/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {
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
		h.logger.Error("failed to refresh token", "error", err)
		c.JSON(http.StatusInternalServerError, auth.ErrorResponse{
			Error:  "failed to refresh token",
			Detail: err.Error(),
		})
		return
	}

	const maxAge = 86400
	const domain = "localhost"

	c.SetCookie(
		"refresh",
		result.RefreshToken,
		maxAge,
		"/",
		domain,
		true,
		true,
	)

	c.JSON(http.StatusOK, result)
}
