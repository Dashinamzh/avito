package user

import (
	"context"

	"github.com/Dashinamzh/avito/internal/model"
	"github.com/google/uuid"
)

func (s *UserService) Get(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user, err := s.UserRepo.Get(ctx, id)
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}
