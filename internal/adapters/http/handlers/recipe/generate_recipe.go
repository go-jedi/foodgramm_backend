package recipe

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// GenerateRecipe
// @Summary Generate a new recipe
// @Description Generates a new recipe based on the provided parameters.
// @Tags Recipe
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body recipe.GenerateRecipeDTO true "Generate recipe request body"
// @Success 200 {object} recipe.Recipes "Successfully generated recipe"
// @Failure 400 {object} recipe.ErrorResponse "Bad request due to invalid input"
// @Failure 404 {object} recipe.ErrorResponse "User does not exist"
// @Failure 500 {object} recipe.ErrorResponse "Internal server error"
// @Router /v1/recipe/generate [post]
func (h *Handler) generateRecipe(c *gin.Context) {
	h.logger.Debug("[generateRecipe] execute handler")

	var dto recipe.GenerateRecipeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	dto.SetDefaults()

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.recipeService.GenerateRecipe(c.Request.Context(), dto)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, recipe.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to generate recipe", "error", err)
		c.JSON(http.StatusInternalServerError, recipe.ErrorResponse{
			Error:  "failed to generate recipe",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
