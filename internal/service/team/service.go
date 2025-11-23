package team

import (
	repo "github.com/Dashinamzh/avito/internal/repo"
	def "github.com/Dashinamzh/avito/internal/service"
)

var _ def.TeamService = (*TeamService)(nil)

type TeamService struct {
	TeamRepo repo.TeamRepository
}

func NewTeamService(TeamRepo repo.TeamRepository) *TeamService {
	return &TeamService{
		TeamRepo: TeamRepo,
	}
}
