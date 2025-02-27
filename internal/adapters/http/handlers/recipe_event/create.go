package recipeevent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
)

// @Summary Create a new recipe event
// @Description Create a new recipe event with the provided details.
// @Tags Recipe event
// @Accept json
// @Produce json
// @Param request body recipeevent.CreateDTO true "Recipe event data"
// @Success 201 {object} recipeevent.Recipe
// @Failure 400 {object} recipeevent.ErrorResponse
// @Failure 409 {object} recipeevent.ErrorResponse
// @Failure 500 {object} recipeevent.ErrorResponse
// @Router /v1/recipe_event [post]
func (h *Handler) create(c *gin.Context) {
	h.logger.Debug("[create a new recipe event] execute handler")

	var dto recipeevent.CreateDTO
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

	result, err := h.recipeEventService.Create(c.Request.Context(), dto)
	if err != nil {
		h.logger.Error("failed to create recipe event", "error", err)
		c.JSON(http.StatusInternalServerError, recipeevent.ErrorResponse{
			Error:  "failed to create recipe event",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}
