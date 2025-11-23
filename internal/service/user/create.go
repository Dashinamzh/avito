package user

import (
	"context"

	"github.com/Dashinamzh/avito/internal/model"
)

func (s *UserService) Create(ctx context.Context, user *model.User) error {
	err := s.UserRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
