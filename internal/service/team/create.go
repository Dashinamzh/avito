package team

import (
	"context"

	"github.com/Dashinamzh/avito/internal/model"
)

func (s *TeamService) Create(ctx context.Context, team *model.Team) error {
	err := s.TeamRepo.Create(ctx, team)
	if err != nil {
		return err
	}
	return nil
}
