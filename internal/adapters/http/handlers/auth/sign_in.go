package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
)

// SignIn
// @Summary Sign in user
// @Description Sign in a user using their Telegram ID, username, first name, and last name
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body auth.SignInDTO true "Sign in request body"
// @Success 200 {object} auth.SignInResp "Successful response with tokens"
// @Failure 400 {object} auth.ErrorResponse "Bad request error"
// @Failure 500 {object} auth.ErrorResponse "Internal server error"
// @Router /v1/auth/signin [post]
func (h *Handler) SignIn(c *gin.Context) {
	var dto auth.SignInDTO
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

	result, err := h.authService.SignIn(c.Request.Context(), dto)
	if err != nil {
		h.logger.Error("failed to sign in user", "error", err)
		c.JSON(http.StatusInternalServerError, auth.ErrorResponse{
			Error:  "failed to sign in user",
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
