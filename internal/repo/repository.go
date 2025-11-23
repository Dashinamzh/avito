package repo

import (
	"context"

	"github.com/Dashinamzh/avito/internal/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id uuid.UUID) (*model.User, error)
	SetActive(ctx context.Context, id uuid.UUID, is_active bool) error
}

type TeamRepository interface {
	Create(ctx context.Context, team *model.Team) error
	Get(ctx context.Context, name string) (*model.Team, error)
	// Delete удаляет команду
	Delete(ctx context.Context, name string) error
}

type PullRequestRepository interface {
	Create(ctx context.Context, pr *model.PullRequest) error
	Get(ctx context.Context, id uuid.UUID) (*model.PullRequest, error)
}
