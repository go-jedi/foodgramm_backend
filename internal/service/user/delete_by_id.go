package user

import "context"

func (s *serv) DeleteByID(ctx context.Context, id int64) (int64, error) {
	s.logger.Debug("[delete by id] execute service")

	return s.userRepository.DeleteByID(ctx, id)
}
