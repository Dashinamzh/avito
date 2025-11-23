package pr

import (
	"context"

	"github.com/Dashinamzh/avito/internal/repo/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (r *prRepo) Create(ctx context.Context, pr *model.PullRequest) error {
	query := `
        INSERT INTO pull_requests 
        (title, author_id, status, reviewers, created_at, merged_at) 
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `

	// Инициализируем reviewers если nil
	if pr.Reviewers == nil {
		pr.Reviewers = []uuid.UUID{}
	}

	err := r.db.QueryRowContext(ctx, query,
		pr.Title,
		pr.AuthorID,
		pr.Status,
		pq.Array(pr.Reviewers), // преобразуем slice в PostgreSQL array
		pr.CreatedAt,
		pr.MergedAt,
	).Scan(&pr.ID)

	return err
}
