package user

import "context"

func (s *serv) GetUserCount(ctx context.Context) (int64, error) {
	return s.userRepository.GetUserCount(ctx)
}
