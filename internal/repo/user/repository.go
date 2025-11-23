package user

import (
	"database/sql"

	"github.com/Dashinamzh/avito/internal/repo"
)

var _ repo.UserRepository = (*userRepo)(nil)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepository {
	return &userRepo{db: db}
}
