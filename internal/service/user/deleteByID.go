package user

import "context"

func (s *serv) DeleteByID(ctx context.Context, id int64) (int64, error) {
	return s.userRepository.DeleteByID(ctx, id)
}
