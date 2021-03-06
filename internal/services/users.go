package services

import (
	"github.com/yvv4git/users-service/internal/domain"
	"github.com/yvv4git/users-service/internal/repository"
)

// UsersService ...
type UsersService struct {
	repo repository.UsersRepository
}

// NewUsersService is used as constructor for UsersService.
func NewUsersService(repo repository.UsersRepository) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

// Create is used for create user in repository.
func (s *UsersService) Create(user domain.Users) (result *domain.Users, err error) {
	return s.repo.Create(user)
}

// Read is used for find user in repository by id.
func (s *UsersService) Read(id int64) (result *domain.Users, err error) {
	return s.repo.Read(id)
}

// Update is used for update user in repository.
func (s *UsersService) Update(user domain.Users) (err error) {
	return s.repo.Update(user)
}

// Del is used for delete user from repository by id.
func (s *UsersService) Del(id int64) (err error) {
	return s.repo.Del(id)
}
