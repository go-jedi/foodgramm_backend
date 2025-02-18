package recipe

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// GetListRecipesByTelegramID
// @Summary Get a list of recipes with pagination by Telegram ID
// @Description Retrieves a paginated list of recipes associated with a specific Telegram ID.
// @Tags Recipe
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body recipe.GetListRecipesByTelegramIDDTO true "Request body containing Telegram ID and pagination details"
// @Success 200 {object} recipe.GetListRecipesByTelegramIDResponseSwagger "Successfully retrieved the list of recipes"
// @Failure 400 {object} recipe.ErrorResponse "Bad request due to invalid input or validation failure"
// @Failure 404 {object} recipe.ErrorResponse "User with the specified Telegram ID does not exist"
// @Failure 500 {object} recipe.ErrorResponse "Internal server error"
// @Router /v1/recipe/list [post]
func (h *Handler) getListRecipesByTelegramID(c *gin.Context) {
	h.logger.Debug("[get list recipes by telegram id] execute handler")

	var dto recipe.GetListRecipesByTelegramIDDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.recipeService.GetListRecipesByTelegramID(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, recipe.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to get list recipes by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, recipe.ErrorResponse{
			Error:  "failed to get list recipes by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
