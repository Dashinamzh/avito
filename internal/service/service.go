package service

import (
	"context"

	"github.com/Dashinamzh/avito/internal/model"
	"github.com/google/uuid"
)

type UserService interface {
	SetActivity(ctx context.Context, userID uuid.UUID, isActive bool) error
	Register(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
}

type TeamService interface {
	Create(ctx context.Context, team *model.Team) error
	Get(ctx context.Context, name string) (*model.Team, error)
}

type PullRequestService interface {
	Create(ctx context.Context, pr *model.PullRequest) error
	Get(ctx context.Context, id uuid.UUID) (*model.PullRequest, error)
}
