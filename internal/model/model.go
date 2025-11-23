package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	IsActive bool      `json:"is_active" db:"is_active"`
}

type Team struct {
	Name    uuid.UUID `json:"id" db:"name"`
	Members []*User   `json:"members"`
}

type PRStatus string

const (
	PRStatusOpen   PRStatus = "OPEN"
	PRStatusMerged PRStatus = "MERGED"
)

type PullRequest struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	Title     string      `json:"title" db:"title"`
	AuthorID  uuid.UUID   `json:"author_id" db:"author_id"`
	Status    PRStatus    `json:"status" db:"status"`
	Reviewers []uuid.UUID `json:"reviewers"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	MergedAt  *time.Time  `json:"merged_at" db:"merged_at"`
}
