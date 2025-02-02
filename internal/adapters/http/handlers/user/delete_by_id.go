package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Delete a user by ID
// @Description Delete a user by their unique identifier.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param userID path int64 true "User ID"
// @Success 200 {integer} int64 "ID of the deleted user"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/id/{userID} [delete]
func (h *Handler) deleteByID(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		h.logger.Error("failed to get param userID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to get param userID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	const (
		base    = 10
		bitSize = 64
	)

	userIDInt, err := strconv.ParseInt(userID, base, bitSize)
	if err != nil {
		h.logger.Error("failed parse string to int64", "error", err)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed parse string to int64",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.userService.DeleteByID(c.Request.Context(), userIDInt)
	if err != nil {
		h.logger.Error("failed to delete user by id", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to delete user by id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
