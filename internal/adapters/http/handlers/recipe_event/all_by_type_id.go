package recipeevent

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Get all recipes event by type id
// @Description Retrieve a list of all recipe's event by type id.
// @Tags Recipe event
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param typeID path int64 true "Type ID"
// @Success 200 {object} recipeevent.Recipe "Recipes event details"
// @Failure 400 {object} recipeevent.ErrorResponse
// @Failure 500 {object} recipeevent.ErrorResponse
// @Router /v1/recipe_event/all/type/{typeID} [get]
func (h *Handler) allByTypeID(c *gin.Context) {
	h.logger.Debug("[get all recipes event by type id] execute handler")

	typeID := c.Param("typeID")
	if typeID == "" {
		h.logger.Error("failed to get param typeID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "failed to get param typeID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	const (
		base    = 10
		bitSize = 64
	)

	typeIDInt, err := strconv.ParseInt(typeID, base, bitSize)
	if err != nil {
		h.logger.Error("failed parse string to int64", "error", err)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "failed parse string to int64",
			Detail: err.Error(),
		})
		return
	}

	if typeIDInt <= 0 {
		h.logger.Error("typeIDInt is less than or equal to zero", "error", err)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "typeIDInt must be greater than zero",
			Detail: "the provided typeIDInt is less than or equal to zero, which is not allowed",
		})
		return
	}

	result, err := h.recipeEventService.AllByTypeID(c.Request.Context(), typeIDInt)
	if err != nil {
		h.logger.Error("failed to get recipes event by type id", "error", err)
		c.JSON(http.StatusInternalServerError, recipeevent.ErrorResponse{
			Error:  "failed to get recipes event by type id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
