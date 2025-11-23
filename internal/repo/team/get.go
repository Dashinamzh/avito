package team

import (
	"context"
	"database/sql"

	"github.com/Dashinamzh/avito/internal/model"
)

func (r *teamRepo) Get(ctx context.Context, name string) (*model.Team, error) {
	teamQuery := `SELECT name FROM teams WHERE name = $1`
	var team model.Team
	err := r.db.QueryRowContext(ctx, teamQuery, name).Scan(&team.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &team, nil
}
