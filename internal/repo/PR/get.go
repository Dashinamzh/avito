package pr

import (
	"context"
	"database/sql"

	"github.com/Dashinamzh/avito/internal/repo/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (r *prRepo) Get(ctx context.Context, id uuid.UUID) (*model.PullRequest, error) {
	query := `
        SELECT id, title, author_id, status, reviewers, created_at, merged_at
        FROM pull_requests 
        WHERE id = $1
    `

	var pr model.PullRequest
	var reviewers pq.StringArray

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&pr.ID,
		&pr.Title,
		&pr.AuthorID,
		&pr.Status,
		&reviewers,
		&pr.CreatedAt,
		&pr.MergedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Конвертируем строки в UUID
	pr.Reviewers = make([]uuid.UUID, len(reviewers))
	for i, reviewerStr := range reviewers {
		reviewerID, err := uuid.Parse(reviewerStr)
		if err != nil {
			return nil, err
		}
		pr.Reviewers[i] = reviewerID
	}

	return &pr, nil
}
