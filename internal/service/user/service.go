package user

import (
	repo "github.com/Dashinamzh/avito/internal/repo"
	def "github.com/Dashinamzh/avito/internal/service"
)

var _ def.UserService = (*UserService)(nil)

type UserService struct {
	UserRepo repo.UserRepository
}

func NewUserService(UserRepo repo.UserRepository) *UserService {
	return &UserService{
		UserRepo: UserRepo,
	}
}
