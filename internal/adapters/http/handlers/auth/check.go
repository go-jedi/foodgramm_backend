package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// Check
// @Summary Check user token
// @Description Check if the provided Telegram ID and token are valid
// @Tags Authentication
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body auth.CheckDTO true "Check request body"
// @Success 200 {object} auth.CheckResponse "Successful response with user token details"
// @Failure 400 {object} auth.ErrorResponse "Bad request error"
// @Failure 500 {object} auth.ErrorResponse "Internal server error"
// @Router /v1/auth/check [post]
func (h *Handler) check(c *gin.Context) {
	h.logger.Debug("[check user token] execute handler")

	var dto auth.CheckDTO
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

	result, err := h.authService.Check(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, auth.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to check user token", "error", err)
		c.JSON(http.StatusInternalServerError, auth.ErrorResponse{
			Error:  "failed to check user token",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
