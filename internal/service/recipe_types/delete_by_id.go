package recipetypes

import "context"

func (s *serv) DeleteByID(ctx context.Context, recipeTypeID int64) (int64, error) {
	return s.recipeTypesRepository.DeleteByID(ctx, recipeTypeID)
}
