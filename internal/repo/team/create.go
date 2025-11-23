package team

import (
	"context"

	"github.com/Dashinamzh/avito/internal/repo/model"
)

func (r *teamRepo) Create(ctx context.Context, team *model.Team) error {
	query := `
		 INSERT INTO teams (id, name) 
        VALUES ($1, $2)
	`
	_, err := r.db.ExecContext(ctx, query, team.Name, team.Members)
	if err != nil {
		return err
	}
	return nil
}
