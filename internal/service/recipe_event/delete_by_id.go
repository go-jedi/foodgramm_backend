package recipeevent

import "context"

func (s *serv) DeleteByID(ctx context.Context, recipeID int64) (int64, error) {
	return s.recipeEventRepository.DeleteByID(ctx, recipeID)
}
