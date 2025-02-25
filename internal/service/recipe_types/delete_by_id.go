package recipetypes

import "context"

func (s *serv) DeleteByID(ctx context.Context, recipeTypeID int64) (int64, error) {
	s.logger.Debug("[delete recipe type by id] execute service")

	return s.recipeTypesRepository.DeleteByID(ctx, recipeTypeID)
}
