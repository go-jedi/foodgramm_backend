package recipeevent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	recipeevent "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_event"
)

// @Summary Update a recipe event
// @Description Update a recipe event information.
// @Tags Recipe event
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param request body recipeevent.UpdateDTO true "Recipe event update data"
// @Success 200 {object} recipeevent.Recipe "Updated recipe event details"
// @Failure 400 {object} recipeevent.ErrorResponse
// @Failure 500 {object} recipeevent.ErrorResponse
// @Router /v1/recipe_types [put]
func (h *Handler) update(c *gin.Context) {
	h.logger.Debug("[update recipe event] execute handler")

	var dto recipeevent.UpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		h.logger.Error("failed to bind body", "error", err)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "failed to bind body",
			Detail: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(dto); err != nil {
		h.logger.Error("failed to validate struct", "error", err)
		c.JSON(http.StatusBadRequest, recipeevent.ErrorResponse{
			Error:  "failed to validate struct",
			Detail: err.Error(),
		})
		return
	}

	result, err := h.recipeEventService.Update(c.Request.Context(), dto)
	if err != nil {
		h.logger.Error("failed to update recipe event", "error", err)
		c.JSON(http.StatusInternalServerError, recipeevent.ErrorResponse{
			Error:  "failed to update recipe event",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
