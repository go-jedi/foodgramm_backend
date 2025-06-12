package userblacklist

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	userblacklist "github.com/go-jedi/foodgramm_backend/internal/domain/user_blacklist"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Block a user
// @Description Block a user with the provided details. Accessible only to administrators.
// @Tags UserBlacklist
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the admin user"
// @Param blockUser body userblacklist.BlockUserDTO true "Block user data"
// @Success 201 {object} userblacklist.UsersBlackList
// @Failure 400 {object} userblacklist.ErrorResponse
// @Failure 401 {object} userblacklist.ErrorResponse
// @Failure 403 {object} userblacklist.ErrorResponse
// @Failure 409 {object} userblacklist.ErrorResponse
// @Failure 500 {object} userblacklist.ErrorResponse
// @Router /v1/user/blacklist/block [post]
func (h *Handler) blockUser(c *gin.Context) {
	h.logger.Debug("[block user] execute handler")

	telegramID, err := h.middleware.Auth.GetTelegramIDFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	var dto userblacklist.BlockUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, userblacklist.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, userblacklist.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.userBlackListService.BlockUser(c.Request.Context(), dto, telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserInBlackListAlreadyExists) {
			h.logger.Error("user in blacklist already exists", "error", err)
			c.JSON(http.StatusConflict, userblacklist.ErrorResponse{
				Error:  "user in blacklist already exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to block user", "error", err)
		c.JSON(http.StatusInternalServerError, userblacklist.ErrorResponse{
			Error:  "failed to block user",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
