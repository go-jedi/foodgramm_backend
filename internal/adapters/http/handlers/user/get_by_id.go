package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Get a user by ID
// @Description Retrieve a user by their unique identifier.
// @Tags users
// @Accept json
// @Produce json
// @Param userID path int64 true "User ID"
// @Success 200 {object} user.User "User details"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/id/{userID} [get]
func (h *Handler) getByID(c *gin.Context) {
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

	if userIDInt <= 0 {
		h.logger.Error("userID is less than or equal to zero", "error", err)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "userID must be greater than zero",
			Detail: "the provided userID is less than or equal to zero, which is not allowed",
		})
		return
	}

	result, err := h.userService.GetByID(c, userIDInt)
	if err != nil {
		h.logger.Error("failed to get user by id", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to get user by id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
