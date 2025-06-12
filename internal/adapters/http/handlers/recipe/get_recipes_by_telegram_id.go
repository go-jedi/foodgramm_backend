package recipe

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/internal/domain/recipe"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

// @Summary Get Recipes by Telegram ID
// @Description Retrieve recipes for a user by their Telegram ID
// @Tags Recipe
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the user"
// @Success 200 {array} recipe.Recipes "List of recipes"
// @Failure 400 {object} recipe.ErrorResponse "Bad request error"
// @Failure 404 {object} recipe.ErrorResponse "User not found error"
// @Failure 500 {object} recipe.ErrorResponse "Internal server error"
// @Router /v1/recipe/telegram/{telegramID} [get]
func (h *Handler) getRecipesByTelegramID(c *gin.Context) {
	h.logger.Debug("[get recipes by telegram id] execute handler")

	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.recipeService.GetRecipesByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, recipe.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to get recipes by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, recipe.ErrorResponse{
			Error:  "failed to get recipes by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
