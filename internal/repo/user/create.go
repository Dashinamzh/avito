package user

import (
	"context"
	"log"

	"github.com/Dashinamzh/avito/internal/model"
)

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	query := ` 
		INSERT INTO users (id, name, is_active) 
        VALUES ($1, $2, $3)
        ON CONFLICT (id) DO UPDATE SET
            is_active = EXCLUDED.is_active
		`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.IsActive)
	if err != nil {
		log.Printf("Query err")
	}
	return err
}
