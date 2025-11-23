package pr

import (
	"database/sql"

	"github.com/Dashinamzh/avito/internal/repo"
)

var _ repo.PullRequestRepository = (*prRepo)(nil)

type prRepo struct {
	db *sql.DB
}

func NewPrRepo(db *sql.DB) repo.PullRequestRepository {
	return &prRepo{db: db}
}
