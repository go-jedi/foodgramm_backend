package user

import "context"

func (s *serv) ExistsByID(ctx context.Context, userID int64) (bool, error) {
	return s.userRepository.ExistsByID(ctx, userID)
}
