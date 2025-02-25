package recipetypes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Delete a recipe type by ID
// @Description Delete a recipe type by their unique identifier.
// @Tags Recipe types
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param recipeTypeID path int64 true "Recipe type ID"
// @Success 200 {integer} int64 "ID of the deleted recipe type"
// @Failure 400 {object} recipetypes.ErrorResponse
// @Failure 500 {object} recipetypes.ErrorResponse
// @Router /v1/recipe_types/id/{recipeTypeID} [delete]
func (h *Handler) deleteByID(c *gin.Context) {
	h.logger.Debug("[delete recipe type by id] execute handler")

	recipeTypeID := c.Param("recipeTypeID")
	if recipeTypeID == "" {
		h.logger.Error("failed to get param recipeTypeID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, recipetypes.ErrorResponse{
			Error:  "failed to get param recipeTypeID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	const (
		base    = 10
		bitSize = 64
	)

	recipeTypeIDInt, err := strconv.ParseInt(recipeTypeID, base, bitSize)
	if err != nil {
		h.logger.Error("failed parse string to int64", "error", err)
		c.JSON(http.StatusBadRequest, recipetypes.ErrorResponse{
			Error:  "failed parse string to int64",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.recipeTypesService.DeleteByID(c.Request.Context(), recipeTypeIDInt)
	if err != nil {
		h.logger.Error("failed to delete recipe type by id", "error", err)
		c.JSON(http.StatusInternalServerError, recipetypes.ErrorResponse{
			Error:  "failed to delete recipe type by id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
