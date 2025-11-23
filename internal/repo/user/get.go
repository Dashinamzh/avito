package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/Dashinamzh/avito/internal/model"
	"github.com/google/uuid"
)

func (r *userRepo) Get(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `SELECT id, name, is_active FROM users WHERE id = $1`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.IsActive,
	)
	if err == sql.ErrNoRows {
		log.Printf("no Rows in db")
		return nil, nil
	}
	return &user, err
}
