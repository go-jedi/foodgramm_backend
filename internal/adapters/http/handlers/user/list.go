package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

// @Summary Get a list of users with pagination
// @Description Retrieve a list of users with pagination based on page and size.
// @Tags User
// @Accept json
// @Produce json
// @Param request body user.ListDTO true "Pagination parameters"
// @Success 200 {object} user.ListResponseSwagger "List of users with pagination details"
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /v1/user/list [post]
func (h *Handler) list(c *gin.Context) {
	var dto user.ListDTO
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

	result, err := h.userService.List(c, dto)
	if err != nil {
		h.logger.Error("failed to get list users", "error", err)
		c.JSON(http.StatusInternalServerError, user.ErrorResponse{
			Error:  "failed to get list users",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
