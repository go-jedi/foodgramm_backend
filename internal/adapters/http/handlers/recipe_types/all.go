package recipetypes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
)

// @Summary Get all recipe types
// @Description Retrieve a list of all recipe types.
// @Tags Recipe types
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Success 200 {array} recipetypes.RecipeTypes
// @Failure 500 {object} recipetypes.ErrorResponse
// @Router /v1/recipe_types/all [get]
func (h *Handler) all(c *gin.Context) {
	h.logger.Debug("[get all recipe types] execute handler")

	result, err := h.recipeTypesService.All(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to get all recipe types", "error", err)
		c.JSON(http.StatusInternalServerError, recipetypes.ErrorResponse{
			Error:  "failed to get all recipe types",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
