package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Create a new user
// @Description Create a new user with the provided details.
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.CreateDTO true "User data"
// @Success 201 {object} user.User
// @Failure 400 {object} user.ErrorResponse
// @Failure 409 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user [post]
func (h *Handler) create(c *gin.Context) {
	var dto user.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, user.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.userService.Create(c, dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserAlreadyExists) {
			h.logger.Error("user already exists", "error", err)
			c.JSON(http.StatusConflict, user.ErrorResponse{
				Error:  "user already exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to create user",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
