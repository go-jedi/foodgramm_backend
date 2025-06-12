package recipetypes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Create a new recipe type
// @Description Create a new recipe type with the provided details.
// @Tags Recipe types
// @Accept json
// @Produce json
// @Param request body recipetypes.CreateDTO true "Recipe type data"
// @Success 201 {object} recipetypes.RecipeTypes
// @Failure 400 {object} recipetypes.ErrorResponse
// @Failure 409 {object} recipetypes.ErrorResponse
// @Failure 500 {object} recipetypes.ErrorResponse
// @Router /v1/recipe_types [post]
func (h *Handler) create(c *gin.Context) {
	h.logger.Debug("[create a new recipe type] execute handler")

	var dto recipetypes.CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, recipetypes.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, recipetypes.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.recipeTypesService.Create(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrRecipeTypeAlreadyExists) {
			h.logger.Error("recipe type already exists", "error", err)
			c.JSON(http.StatusConflict, recipetypes.ErrorResponse{
				Error:  "recipe type already exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to create recipe type", "error", err)
		c.JSON(http.StatusInternalServerError, recipetypes.ErrorResponse{
			Error:  "failed to create recipe type",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
