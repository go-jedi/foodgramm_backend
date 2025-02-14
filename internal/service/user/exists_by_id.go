package user

import "context"

func (s *serv) ExistsByID(ctx context.Context, userID int64) (bool, error) {
	s.logger.Debug("[ExistsByID] execute service")

	return s.userRepository.ExistsByID(ctx, userID)
}
