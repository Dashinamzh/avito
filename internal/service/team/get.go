package team

import (
	"context"

	"github.com/Dashinamzh/avito/internal/model"
)

func (s *TeamService) Get(ctx context.Context, name string) (*model.Team, error) {
	team, err := s.TeamRepo.Get(ctx, name)
	if err != nil {
		return &model.Team{}, err
	}

	return team, nil
}
