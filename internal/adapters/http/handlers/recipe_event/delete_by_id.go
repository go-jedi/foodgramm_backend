package recipeevent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Delete a recipe event by ID
// @Description Delete a recipe event by their unique identifier.
// @Tags Recipe event
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param recipeID path int64 true "Recipe event ID"
// @Success 200 {integer} int64 "ID of the deleted recipe event"
// @Failure 400 {object} recipeevent.ErrorResponse
// @Failure 500 {object} recipeevent.ErrorResponse
// @Router /v1/recipe_event/id/{recipeID} [delete]
func (h *Handler) deleteByID(c *gin.Context) {
	h.logger.Debug("[delete recipe event by id] execute handler")

	recipeID := c.Param("recipeID")
	if recipeID == "" {
		h.logger.Error("failed to get param recipeID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "failed to get param recipeID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	const (
		base    = 10
		bitSize = 64
	)

	recipeIDInt, err := strconv.ParseInt(recipeID, base, bitSize)
	if err != nil {
		h.logger.Error("failed parse string to int64", "error", err)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "failed parse string to int64",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.recipeEventService.DeleteByID(c.Request.Context(), recipeIDInt)
	if err != nil {
		h.logger.Error("failed to delete recipe event by id", "error", err)
		c.JSON(http.StatusInternalServerError, recipeevent.ErrorResponse{
			Error:  "failed to delete recipe event by id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
