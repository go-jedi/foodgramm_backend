package user

import "context"

func (s *serv) ExistsByID(ctx context.Context, userID int64) (bool, error) {
	s.logger.Debug("[check user exists by id] execute service")

	return s.userRepository.ExistsByID(ctx, userID)
}
