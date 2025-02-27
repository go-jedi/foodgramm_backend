package recipeevent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Get a recipe event by ID
// @Description Retrieve a recipe event by their unique identifier.
// @Tags Recipe event
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param recipeID path int64 true "Recipe event ID"
// @Success 200 {object} recipeevent.Recipe "Recipe event details"
// @Failure 400 {object} recipeevent.ErrorResponse
// @Failure 500 {object} recipeevent.ErrorResponse
// @Router /v1/recipe_types/id/{recipeTypeID} [get]
func (h *Handler) getByID(c *gin.Context) {
	h.logger.Debug("[get recipe event by id] execute handler")

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

	if recipeIDInt <= 0 {
		h.logger.Error("recipeID is less than or equal to zero", "error", err)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "recipeID must be greater than zero",
			Detail: "the provided recipeID is less than or equal to zero, which is not allowed",
		})
		return
	}

	result, err := h.recipeEventService.GetByID(c.Request.Context(), recipeIDInt)
	if err != nil {
		h.logger.Error("failed to get recipe event by id", "error", err)
		c.JSON(http.StatusInternalServerError, recipeevent.ErrorResponse{
			Error:  "failed to get recipe event by id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
