package team

import (
	"database/sql"

	"github.com/Dashinamzh/avito/internal/repo"
)

var _ repo.TeamRepository = (*teamRepo)(nil)

type teamRepo struct {
	db *sql.DB
}

func NewTeamRepo(db *sql.DB) repo.TeamRepository {
	return &teamRepo{db: db}
}
