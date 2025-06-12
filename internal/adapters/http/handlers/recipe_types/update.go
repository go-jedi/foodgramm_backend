package recipetypes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
)

// @Summary Update a recipe type
// @Description Update a recipe type information.
// @Tags Recipe types
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body recipetypes.UpdateDTO true "Recipe type update data"
// @Success 200 {object} recipetypes.RecipeTypes "Updated recipe type details"
// @Failure 400 {object} recipetypes.ErrorResponse
// @Failure 500 {object} recipetypes.ErrorResponse
// @Router /v1/recipe_types [put]
func (h *Handler) update(c *gin.Context) {
	h.logger.Debug("[update recipe type] execute handler")

	var dto recipetypes.UpdateDTO
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

	result, err := h.recipeTypesService.Update(c.Request.Context(), dto)
	if err != nil {
		h.logger.Error("failed to update recipe type", "error", err)
		c.JSON(http.StatusInternalServerError, recipetypes.ErrorResponse{
			Error:  "failed to update recipe type",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
