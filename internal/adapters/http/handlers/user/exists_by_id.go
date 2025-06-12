package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Check if a user exists by ID
// @Description Check if a user exists by their unique identifier.
// @Tags User
// @Accept json
// @Produce json
// @Param userID path int64 true "User ID"
// @Success 200 {boolean} boolean "Boolean flag indicating if the user exists"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/exists/id/{userID} [get]
func (h *Handler) existsByID(c *gin.Context) {
	h.logger.Debug("[check user exists by id] execute handler")

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

	result, err := h.userService.ExistsByID(c.Request.Context(), userIDInt)
	if err != nil {
		h.logger.Error("failed to exists user by id", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to exists user by id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
