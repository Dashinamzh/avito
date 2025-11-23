package user

import (
	"context"

	"github.com/google/uuid"
)

func (r *userRepo) SetActive(ctx context.Context, id uuid.UUID, is_active bool) error {
	query := `UPDATE users SET is_active = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, is_active, id)
	return err
}
