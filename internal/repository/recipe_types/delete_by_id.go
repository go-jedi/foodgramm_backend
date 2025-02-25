package recipetypes

import "context"

func (r *repo) DeleteByID(_ context.Context, _ int64) (int64, error) {
	return 0, nil
}
